package content

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"mistlib/internal/content/ipfs"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

func handleGetWorldData(args []string) {
	headData, err := localFS.ReadFile(".fw/HEAD")
	if err != nil {
		fmt.Printf("[World] Error reading HEAD: %v\n", err)
		return
	}

	var headYamlData HeadData
	err = yaml.Unmarshal(headData, &headYamlData)
	if err != nil {
		fmt.Printf("[World] Error parsing HEAD: %v\n", err)
		return
	}

	worldDataStr, err := localFS.ReadFile(filepath.Join(".fw", "worlds", headYamlData.Guid))
	if err != nil {
		fmt.Printf("[World] Error reading world data: %v\n", err)
		return
	}

	var worldYamlData WorldData
	err = yaml.Unmarshal(worldDataStr, &worldYamlData)
	if err != nil {
		fmt.Printf("[World] Error parsing world data: %v\n", err)
		return
	}

	err = loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	worldAllData := make([]MetaData, 0)
	for _, cid := range worldYamlData.CID {
		metaFilePath := filepath.Join(".fw", "objects", cid)
		file, err := localFS.ReadFile(metaFilePath)
		if err != nil {
			fmt.Printf("[World] Error reading file: %v\n", err)
			return
		}

		decryptedData, err := decryptData(file)
		if err != nil {
			fmt.Printf("[World] Error decrypting world data: %v\n", err)
			return
		}

		decompressedData, err := decompressData(decryptedData)
		if err != nil {
			fmt.Printf("[World] Error decompressing world data: %v\n", err)
			return
		}
		var worldSingleData MetaData
		err = yaml.Unmarshal(decompressedData, &worldSingleData)
		if err != nil {
			fmt.Printf("[World] Error parsing world data: %v\n", err)
			return
		}
		worldAllData = append(worldAllData, worldSingleData)
	}

	savePath := filepath.Join(".fw", "content", headYamlData.CurrentWorld+".yaml")
	worldAllDataYaml, err := yaml.Marshal(&worldAllData)
	if err != nil {
		fmt.Printf("[World] Error parsing world data: %v\n", err)
		return
	}
	err = localFS.WriteFile(savePath, worldAllDataYaml)
	if err != nil {
		fmt.Printf("[World] Error saving world data: %v\n", err)
		return
	}
}

func handleDownloadWorld(args []string) {
	if len(args) < 1 {
		fmt.Println("[World] Usage: fw download-world <cid>")
		return
	}

	cid := args[0]
	worldPath := filepath.Join(".fw", "objects", cid)
	// Simulate IPFS Download
	// err := ipfs.Download(cid, worldPath)
	err := ipfs.Download(cid, worldPath)
	if err != nil {
		fmt.Printf("[World] Error downloading world data: %v\n", err)
		return
	}

	err = loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	worldData, err := localFS.ReadFile(worldPath)
	if err != nil {
		fmt.Printf("[World] Error reading world data: %v\n", err)
		return
	}

	decryptedData, err := decryptData(worldData)
	if err != nil {
		fmt.Printf("[World] Error decrypting world data: %v\n", err)
		return
	}

	decompressedData, err := decompressData(decryptedData)
	if err != nil {
		fmt.Printf("[World] Error decompressing world data: %v\n", err)
		return
	}

	var worldYamlData WorldData
	err = yaml.Unmarshal(decompressedData, &worldYamlData)
	if err != nil {
		fmt.Printf("[World] Error parsing world data: %v\n", err)
		return
	}

	newWorldPath := filepath.Join(".fw", "worlds", worldYamlData.GUID)

	err = localFS.WriteFile(newWorldPath, decompressedData)
	if err != nil {
		fmt.Printf("[World] Error saving world data: %v\n", err)
		return
	}
}

func handleGetWorldCID(args []string) {
	headData, err := localFS.ReadFile(".fw/HEAD")
	if err != nil {
		fmt.Printf("[World] Error reading HEAD: %v\n", err)
		return
	}

	var headYamlData HeadData
	err = yaml.Unmarshal(headData, &headYamlData)
	if err != nil {
		fmt.Printf("[World] Error parsing HEAD: %v\n", err)
		return
	}
	guid := headYamlData.Guid
	filePath := filepath.Join(".fw", "worlds", guid)

	file, err := localFS.ReadFile(filePath)
	if err != nil {
		fmt.Printf("[World] Error opening file %s: %v\n", filePath, err)
		return
	}

	hash := sha256.Sum256(file)
	hashString := hex.EncodeToString(hash[:])

	compressedData, err := compressData(file)
	if err != nil {
		fmt.Printf("[World] Error compressing file: %v\n", err)
		return
	}

	err = loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	encryptedData, err := encryptData(compressedData)
	if err != nil {
		fmt.Printf("[World] Error encrypting file: %v\n", err)
		return
	}

	objectPath := filepath.Join(".fw", "objects", hashString)
	err = localFS.WriteFile(objectPath, encryptedData)
	if err != nil {
		fmt.Printf("[World] Error saving file %s: %v\n", objectPath, err)
		return
	}

	// Simulate IPFS Upload
	cid, err := ipfs.Upload(objectPath)
	if err != nil {
		fmt.Printf("[World] Error uploading file to IPFS: %v\n", err)
		return
	}

	fmt.Printf("[World] World data uploaded to IPFS with CID: %s\n", cid)

	// ファイル名を変更
	err = localFS.Rename(objectPath, filepath.Join(".fw", "objects", cid))
	if err != nil {
		fmt.Printf("[World] Error renaming file: %v\n", err)
		return
	}
}

func handleSwitch(args []string) {
	if len(args) < 1 {
		fmt.Println("[World] Usage: fw create <world-name>")
		return
	}

	worldName := args[0]

	fmt.Println("[World] Creating new world...")

	worldListPath := filepath.Join(".fw", "world_list")
	worldDict, err := loadWorldList(worldListPath)
	if err != nil {
		fmt.Printf("[World] Error reading world list: %v\n", err)
		return
	}

	if _, exists := worldDict[worldName]; exists {
		fmt.Printf("[World] Error: World %s already exists\n", worldName)
		return
	}

	guid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("[World] Error generating GUID: %v\n", err)
		return
	}

	worldGuid := guid.String()
	err = updateHead(worldName, worldGuid)
	if err != nil {
		fmt.Printf("[World] Error updating HEAD: %v\n", err)
		return
	}

	err = createWorldFile(worldName, worldGuid)
	if err != nil {
		fmt.Printf("[World] Error creating world file: %v\n", err)
		return
	}

	worldDict[worldName] = worldGuid
	err = saveWorldList(worldListPath, worldDict)
	if err != nil {
		fmt.Printf("[World] Error updating world list: %v\n", err)
		return
	}
}

func loadWorldList(worldListPath string) (map[string]string, error) {
	worldListData, err := localFS.ReadFile(worldListPath)
	if err != nil {
		return nil, err
	}

	worldList := strings.Split(string(worldListData), "\n")
	worldDict := make(map[string]string)
	for _, line := range worldList {
		if line == "" {
			continue
		}
		kv := strings.Split(line, ": ")
		worldDict[kv[0]] = kv[1]
	}
	return worldDict, nil
}

func updateHead(worldName, worldGuid string) error {
	headPath := filepath.Join(".fw", "HEAD")
	headContent := fmt.Sprintf("currentWorld: %s\nguid: %s\n", worldName, worldGuid)
	return localFS.WriteFile(headPath, []byte(headContent))
}

func createWorldFile(worldName, worldGuid string) error {
	worldPath := filepath.Join(".fw", "worlds", worldGuid)
	worldMeta := fmt.Sprintf("name: %s\nguid: %s\ncid: %s\n", worldName, worldGuid, "")
	return localFS.WriteFile(worldPath, []byte(worldMeta))
}

func saveWorldList(worldListPath string, worldDict map[string]string) error {
	var newWorldData []byte
	for k, v := range worldDict {
		newWorldData = append(newWorldData, []byte(fmt.Sprintf("%s: %s\n", k, v))...)
	}
	return localFS.WriteFile(worldListPath, newWorldData)
}

func handleGet(args []string) {
	if len(args) < 1 {
		fmt.Println("[World] Usage: fw get <cid>")
		return
	}

	err := loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	cid := args[0]
	metaDataPath := filepath.Join(".fw", "objects", cid)

	// Simulate IPFS Download
	err = ipfs.Download(cid, metaDataPath)
	if err != nil {
		fmt.Printf("[World] Error downloading metadata: %v\n", err)
		return
	}

	encryptedMetaData, err := localFS.ReadFile(metaDataPath)
	if err != nil {
		fmt.Printf("[World] Error reading metadata %s: %v\n", metaDataPath, err)
		return
	}

	decryptedMetaData, err := decryptData(encryptedMetaData)
	if err != nil {
		fmt.Printf("[World] Error decrypting metadata %s: %v\n", metaDataPath, err)
		return
	}

	decompressedMetaData, err := decompressData(decryptedMetaData)
	if err != nil {
		fmt.Printf("[World] Error decompressing metadata %s: %v\n", metaDataPath, err)
		return
	}

	metaDataContent := string(decompressedMetaData)
	fileCid, fileName, err := extractFileDetailsFromMetaData(metaDataContent)
	if err != nil {
		fmt.Printf("[World] Error extracting file details from metadata %s: %v\n", metaDataPath, err)
		return
	}

	filePath := filepath.Join(".fw", "content", filepath.Base(fileName))
	// Simulate IPFS Download
	err = ipfs.Download(fileCid, filePath)
	if err != nil {
		fmt.Printf("[World] Error downloading file from IPFS: %v\n", err)
		return
	}

	encryptedFileData, err := localFS.ReadFile(filePath)
	if err != nil {
		fmt.Printf("[World] Error reading file %s: %v\n", filePath, err)
		return
	}

	decryptedFileData, err := decryptData(encryptedFileData)
	if err != nil {
		fmt.Printf("[World] Error decrypting file %s: %v\n", filePath, err)
		return
	}

	decompressedFileData, err := decompressData(decryptedFileData)
	if err != nil {
		fmt.Printf("[World] Error decompressing file %s: %v\n", filePath, err)
		return
	}

	err = localFS.WriteFile(filePath, decompressedFileData)
	if err != nil {
		fmt.Printf("[World] Error saving decompressed file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("[World] File downloaded, decrypted, and saved to %s\n", filePath)
}

func handlePut(args []string) {
	if len(args) < 10 {
		fmt.Println("[World] Usage: fw put <file> <x> <y> <z> <rx> <ry> <rz> <sx> <sy> <sz>")
		return
	}

	if err := loadPassword(); err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	filePath, err := filepath.Abs(args[0])
	if err != nil {
		fmt.Printf("[World] Invalid file path: %v\n", err)
		return
	}

	coords, err := parseCoordinates(args[1:])
	if err != nil {
		fmt.Printf("[World] Invalid input: %v\n", err)
		return
	}

	fmt.Printf("[World] Putting world binary data from %s with coordinates (%f, %f, %f), rotation (%f, %f, %f), scale (%f, %f, %f)\n",
		filePath, coords[0], coords[1], coords[2], coords[3], coords[4], coords[5], coords[6], coords[7], coords[8])

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("[World] Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Printf("[World] Error hashing file: %v\n", err)
		return
	}
	file.Seek(0, 0)
	hashString := hex.EncodeToString(hash.Sum(nil))

	byteArray, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("[World] Error reading file: %v\n", err)
		return
	}

	compressedData, err := compressData(byteArray)
	if err != nil {
		fmt.Printf("[World] Error compressing file: %v\n", err)
		return
	}

	encryptedData, err := encryptData(compressedData)
	if err != nil {
		fmt.Printf("[World] Error encrypting file: %v\n", err)
		return
	}

	objectPath := filepath.Join(".fw", "objects", hashString)
	err = localFS.WriteFile(objectPath, encryptedData)
	if err != nil {
		fmt.Printf("[World] Error saving file %s: %v\n", objectPath, err)
		return
	}

	// Simulate IPFS Upload
	cid, err := ipfs.Upload(objectPath)
	if err != nil {
		fmt.Printf("[World] Error uploading file to IPFS: %v\n", err)
		return
	}

	err = localFS.Rename(objectPath, filepath.Join(".fw", "objects", cid))
	if err != nil {
		fmt.Printf("[World] Error renaming file: %v\n", err)
		return
	}

	metaDataContent := generateMetaDataContent(filePath, cid, coords)
	encryptedMetaData, err := encryptAndCompressMetaData(metaDataContent)
	if err != nil {
		fmt.Printf("[World] Error processing metadata: %v\n", err)
		return
	}

	metaDataPath := filepath.Join(".fw", "objects", hashString)
	err = localFS.WriteFile(metaDataPath, encryptedMetaData)
	if err != nil {
		fmt.Printf("[World] Error creating metadata file %s: %v\n", metaDataPath, err)
		return
	}

	metaCid, err := ipfs.Upload(metaDataPath)
	if err != nil {
		fmt.Printf("[World] Error uploading metadata to IPFS: %v\n", err)
		return
	}

	err = localFS.Rename(metaDataPath, filepath.Join(".fw", "objects", metaCid))
	if err != nil {
		fmt.Printf("[World] Error renaming metadata file: %v\n", err)
		return
	}

	fmt.Printf("[World] File saved as %s\n", filepath.Join(".fw", "objects", cid))
	fmt.Printf("[World] Metadata saved as %s\n", filepath.Join(".fw", "objects", metaCid))

	err = updateWorldData(metaCid)
	if err != nil {
		fmt.Printf("[World] Error updating world data: %v\n", err)
		return
	}

	fmt.Println("[World] World data updated successfully.")
}

func handleUpdate(args []string) {
	if len(args) < 10 {
		fmt.Println("[World] Usage: fw update <meta-cid> <x> <y> <z> <rx> <ry> <rz> <sx> <sy> <sz>")
		return
	}

	metaCid := args[0]
	coords, err := parseCoordinates(args[1:])
	if err != nil {
		fmt.Printf("[World] Invalid input: %v\n", err)
		return
	}

	metaFilePath := filepath.Join(".fw", "objects", metaCid)
	metaFile, err := localFS.ReadFile(metaFilePath)
	if err != nil {
		fmt.Printf("[World] Error reading file %s: %v\n", metaFilePath, err)
		return
	}

	// パスワードを読み込む
	err = loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	decryptedMetaFile, err := decryptData(metaFile)
	if err != nil {
		fmt.Printf("[World] Error decrypting file %s: %v\n", metaFilePath, err)
		return
	}

	decompressedMetaFile, err := decompressData(decryptedMetaFile)
	if err != nil {
		fmt.Printf("[World] Error decompressing file %s: %v\n", metaFilePath, err)
		return
	}

	var metaFileData MetaData
	err = yaml.Unmarshal(decompressedMetaFile, &metaFileData)
	if err != nil {
		fmt.Printf("[World] Error parsing file %s: %v\n", metaFilePath, err)
		return
	}

	metaFileData.X = coords[0]
	metaFileData.Y = coords[1]
	metaFileData.Z = coords[2]
	metaFileData.RX = coords[3]
	metaFileData.RY = coords[4]
	metaFileData.RZ = coords[5]
	metaFileData.SX = coords[6]
	metaFileData.SY = coords[7]
	metaFileData.SZ = coords[8]

	newMetaFileData, err := yaml.Marshal(&metaFileData)
	if err != nil {
		fmt.Printf("[World] Error marshalling new metadata: %v\n", err)
		return
	}

	encryptedMetaFileData, err := encryptAndCompressMetaData(newMetaFileData)
	if err != nil {
		fmt.Printf("[World] Error encrypting new metadata: %v\n", err)
		return
	}

	guid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("[World] Error generating GUID: %v\n", err)
		return
	}

	guidString := guid.String()

	newMetaFilePath := filepath.Join(".fw", "objects", guidString)
	err = localFS.WriteFile(newMetaFilePath, encryptedMetaFileData)
	if err != nil {
		fmt.Printf("[World] Error saving new metadata: %v\n", err)
		return
	}

	// ipfs にアップロード
	newMetaCid, err := ipfs.Upload(newMetaFilePath)
	if err != nil {
		fmt.Printf("[World] Error uploading new metadata to IPFS: %v\n", err)
		return
	}

	err = localFS.Rename(newMetaFilePath, filepath.Join(".fw", "objects", newMetaCid))
	if err != nil {
		fmt.Printf("[World] Error renaming new metadata: %v\n", err)
		return
	}

	err = localFS.Remove(metaFilePath)
	if err != nil {
		fmt.Printf("[World] Error deleting old metadata: %v\n", err)
		return
	}

	fmt.Println("[World] World data updated successfully. CID:", newMetaCid)
}

func handleSetCustomData(args []string) {
	if len(args) < 2 {
		fmt.Println("[World] Usage: fw set-custom-data <meta-cid> <value>")
		return
	}

	cid := args[0]
	customData := args[1]
	delete := true // 前のデータを自動で削除するかどうか
	if len(args) > 2 {
		delete = args[2] == "true"
	}

	metaFilePath := filepath.Join(".fw", "objects", cid)
	metaFile, err := localFS.ReadFile(metaFilePath)
	if err != nil {
		fmt.Printf("[World] Error reading file %s: %v\n", metaFilePath, err)
		return
	}

	// パスワードを読み込む
	err = loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	decryptedMetaFile, err := decryptData(metaFile)
	if err != nil {
		fmt.Printf("[World] Error decrypting file %s: %v\n", metaFilePath, err)
		return
	}

	decompressedMetaFile, err := decompressData(decryptedMetaFile)
	if err != nil {
		fmt.Printf("[World] Error decompressing file %s: %v\n", metaFilePath, err)
		return
	}

	var metaFileData MetaData
	err = yaml.Unmarshal(decompressedMetaFile, &metaFileData)
	if err != nil {
		fmt.Printf("[World] Error parsing file %s: %v\n", metaFilePath, err)
		return
	}

	metaFileData.Custom = customData

	newMetaFileData, err := yaml.Marshal(&metaFileData)
	if err != nil {
		fmt.Printf("[World] Error marshalling new metadata: %v\n", err)
		return
	}

	encryptedMetaFileData, err := encryptAndCompressMetaData(newMetaFileData)
	if err != nil {
		fmt.Printf("[World] Error encrypting new metadata: %v\n", err)
		return
	}

	guid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("[World] Error generating GUID: %v\n", err)
		return
	}

	guidString := guid.String()

	newMetaFilePath := filepath.Join(".fw", "objects", guidString)
	err = localFS.WriteFile(newMetaFilePath, encryptedMetaFileData)
	if err != nil {
		fmt.Printf("[World] Error saving new metadata: %v\n", err)
		return
	}

	// ipfs にアップロード
	newMetaCid, err := ipfs.Upload(newMetaFilePath)
	if err != nil {
		fmt.Printf("[World] Error uploading new metadata to IPFS: %v\n", err)
		return
	}

	err = localFS.Rename(newMetaFilePath, filepath.Join(".fw", "objects", newMetaCid))
	if err != nil {
		fmt.Printf("[World] Error renaming new metadata: %v\n", err)
		return
	}

	newMetaFilePath = filepath.Join(".fw", "objects", newMetaCid)

	if delete {
		err = localFS.Remove(metaFilePath)
		if err != nil {
			fmt.Printf("[World] Error deleting old metadata: %v\n", err)
			return
		}
	}

	fmt.Println("[World] Custom data set successfully. New metadata saved as", newMetaFilePath)
}

func handleSetParent(args []string) {
	if len(args) < 2 {
		fmt.Println("[World] Usage: fw set-parent <child-cid> <parent-cid>")
		return
	}

	childCid := args[0]
	parentCid := args[1]
	delete := true // 前のデータを自動で削除するかどうか
	if len(args) > 2 {
		delete = args[2] == "true"
	}

	// load password
	err := loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	parentMetaPath := filepath.Join(".fw", "objects", parentCid)
	parentMetaFile, err := localFS.ReadFile(parentMetaPath)
	if err != nil {
		fmt.Printf("[World] Error reading parent metadata: %v\n", err)
		return
	}

	decryptedParentMetaFile, err := decryptData(parentMetaFile)
	if err != nil {
		fmt.Printf("[World] Error decrypting parent metadata: %v\n", err)
		return
	}

	decompressedParentMetaFile, err := decompressData(decryptedParentMetaFile)
	if err != nil {
		fmt.Printf("[World] Error decompressing parent metadata: %v\n", err)
		return
	}

	var parentMetaData MetaData
	err = yaml.Unmarshal(decompressedParentMetaFile, &parentMetaData)
	if err != nil {
		fmt.Printf("[World] Error parsing parent metadata: %v\n", err)
		return
	}

	parentMetaData.ChildCIDs = append(parentMetaData.ChildCIDs, childCid)

	newParentMetaFile, err := yaml.Marshal(&parentMetaData)
	if err != nil {
		fmt.Printf("[World] Error marshalling new parent metadata: %v\n", err)
		return
	}

	encryptedParentMetaFile, err := encryptAndCompressMetaData(newParentMetaFile)
	if err != nil {
		fmt.Printf("[World] Error encrypting new parent metadata: %v\n", err)
		return
	}

	guid, err := uuid.NewUUID()
	if err != nil {
		fmt.Printf("[World] Error generating GUID: %v\n", err)
		return
	}

	guidString := guid.String()

	newParentMetaPath := filepath.Join(".fw", "objects", guidString)
	err = localFS.WriteFile(newParentMetaPath, encryptedParentMetaFile)
	if err != nil {
		fmt.Printf("[World] Error saving new parent metadata: %v\n", err)
		return
	}

	// ipfs にアップロード
	newParentMetaCid, err := ipfs.Upload(newParentMetaPath)
	if err != nil {
		fmt.Printf("[World] Error uploading new parent metadata to IPFS: %v\n", err)
		return
	}

	err = localFS.Rename(newParentMetaPath, filepath.Join(".fw", "objects", newParentMetaCid))
	if err != nil {
		fmt.Printf("[World] Error renaming new parent metadata: %v\n", err)
		return
	}

	newParentMetaPath = filepath.Join(".fw", "objects", newParentMetaCid)

	if delete {
		err = localFS.Remove(parentMetaPath)
		if err != nil {
			fmt.Printf("[World] Error deleting old parent metadata: %v\n", err)
			return
		}
	}

	fmt.Println("[World] Parent data set successfully. New metadata saved as", newParentMetaPath)
}

func handleSetPassword(args []string) {
	if len(args) < 1 {
		fmt.Println("[World] Usage: fw set-password <password>")
		return
	}

	password := args[0]
	hashedPassword := sha256.Sum256([]byte(password))
	key = hashedPassword[:]

	err := localFS.WriteFile(".fw/password", []byte(hex.EncodeToString(key)))
	if err != nil {
		fmt.Printf("[World] Error saving password: %v\n", err)
		return
	}

	fmt.Println("[World] Password set successfully.")
}

func updateWorldData(metaCid string) error {
	headData, err := localFS.ReadFile(".fw/HEAD")
	if err != nil {
		return err
	}

	var headYamlData HeadData
	err = yaml.Unmarshal(headData, &headYamlData)
	if err != nil {
		return err
	}

	worldDataPath := filepath.Join(".fw", "worlds", headYamlData.Guid)
	worldDataStr, err := localFS.ReadFile(worldDataPath)
	if err != nil {
		return err
	}

	var worldData WorldData
	err = yaml.Unmarshal(worldDataStr, &worldData)
	if err != nil {
		return err
	}

	worldData.CID = append(worldData.CID, metaCid)

	newWorldData, err := yaml.Marshal(&worldData)
	if err != nil {
		return err
	}

	return localFS.WriteFile(worldDataPath, newWorldData)
}

func handleInit(args []string) {
	if _, err := localFS.Stat(".fw"); err == nil {
		fmt.Println("[World] .fw repository already exists.")
		ipfs.InitIPFS()
		return
	}

	fmt.Println("[World] .fw repository initialized.")

	directories := []string{
		".fw",
		".fw/objects",
		".fw/worlds",
		".fw/worlds/heads",
		".fw/worlds/tags",
		".fw/content",
	}

	files := map[string]string{
		".fw/HEAD":        "",
		".fw/world_list":  "",
		".fw/config":      "[core]\n\trepositoryformatversion = 0\n\tfilemode = true\n\tbare = false\n",
		".fw/description": "Unnamed repository; edit this file 'description' to name the repository.\n",
	}

	for _, dir := range directories {
		localFS.MkdirAll(dir)
	}
	for file, content := range files {
		localFS.WriteFile(file, []byte(content))
	}

	ipfs.InitIPFS()
}

func handleCat(args []string) {
	if len(args) < 1 {
		fmt.Println("[World] Usage: fw cat <path>")
		return
	}

	err := loadPassword()
	if err != nil {
		fmt.Printf("[World] Error loading password: %v\n", err)
		return
	}

	// filePath := args[0]
	filePath := filepath.Join(".fw", "objects", args[0])

	encryptedFileData, err := localFS.ReadFile(filePath)
	if err != nil {
		fmt.Printf("[World] Error reading file %s: %v\n", filePath, err)
		return
	}

	decryptedFileData, err := decryptData(encryptedFileData)
	if err != nil {
		fmt.Printf("[World] Error decrypting file %s: %v\n", filePath, err)
		return
	}

	decompressedFileData, err := decompressData(decryptedFileData)
	if err != nil {
		fmt.Printf("[World] Error decompressing file %s: %v\n", filePath, err)
		return
	}

	fmt.Printf("[World] Content of %s:\n%s\n", filePath, string(decompressedFileData))
}
