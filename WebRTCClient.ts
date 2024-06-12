// File: ./WebRTCClient.ts
export class WebRTCClient {
    private peerConnection: RTCPeerConnection;
    private dataChannel: RTCDataChannel;
    private onMessageCallback: (message: any) => void;
    private onConnectionStateChangeCallback: (state: RTCPeerConnectionState) => void;
  
    constructor(iceServers: RTCIceServer[], onMessageCallback: (message: any) => void, onConnectionStateChangeCallback: (state: RTCPeerConnectionState) => void) {
      this.peerConnection = new RTCPeerConnection({ iceServers });
      this.onMessageCallback = onMessageCallback;
      this.onConnectionStateChangeCallback = onConnectionStateChangeCallback;
  
      this.peerConnection.onicecandidate = this.handleICECandidateEvent.bind(this);
      this.peerConnection.oniceconnectionstatechange = this.handleICEConnectionStateChangeEvent.bind(this);
      this.peerConnection.ondatachannel = this.handleDataChannelEvent.bind(this);
    }
  
    public createOffer() {
      this.dataChannel = this.peerConnection.createDataChannel("dataChannel");
      this.setupDataChannel();
      
      this.peerConnection.createOffer()
        .then(offer => this.peerConnection.setLocalDescription(offer))
        .then(() => {
          console.log("Created offer:", this.peerConnection.localDescription);
        })
        .catch(error => console.error("Error creating offer:", error));
    }
  
    public createAnswer(offer: RTCSessionDescriptionInit) {
      this.peerConnection.setRemoteDescription(new RTCSessionDescription(offer))
        .then(() => this.peerConnection.createAnswer())
        .then(answer => this.peerConnection.setLocalDescription(answer))
        .then(() => {
          console.log("Created answer:", this.peerConnection.localDescription);
        })
        .catch(error => console.error("Error creating answer:", error));
    }
  
    public setRemoteDescription(description: RTCSessionDescriptionInit) {
      this.peerConnection.setRemoteDescription(new RTCSessionDescription(description))
        .then(() => console.log("Set remote description"))
        .catch(error => console.error("Error setting remote description:", error));
    }
  
    public addIceCandidate(candidate: RTCIceCandidateInit) {
      this.peerConnection.addIceCandidate(new RTCIceCandidate(candidate))
        .then(() => console.log("Added ICE candidate"))
        .catch(error => console.error("Error adding ICE candidate:", error));
    }
  
    private handleICECandidateEvent(event: RTCPeerConnectionIceEvent) {
      if (event.candidate) {
        console.log("ICE candidate:", event.candidate);
        // Send the ICE candidate to the remote peer
      }
    }
  
    private handleICEConnectionStateChangeEvent() {
      console.log("ICE connection state change:", this.peerConnection.iceConnectionState);
      this.onConnectionStateChangeCallback(this.peerConnection.iceConnectionState);
    }
  
    private handleDataChannelEvent(event: RTCDataChannelEvent) {
      console.log("Data channel event:", event);
      this.dataChannel = event.channel;
      this.setupDataChannel();
    }
  
    private setupDataChannel() {
      this.dataChannel.onopen = () => console.log("Data channel open");
      this.dataChannel.onclose = () => console.log("Data channel closed");
      this.dataChannel.onmessage = (event) => {
        console.log("Data channel message:", event.data);
        this.onMessageCallback(JSON.parse(event.data));
      };
    }
  
    public sendMessage(type: string, data: any) {
      if (this.dataChannel.readyState === "open") {
        this.dataChannel.send(JSON.stringify({ type, data }));
      } else {
        console.error("Data channel is not open");
      }
    }
  }
  