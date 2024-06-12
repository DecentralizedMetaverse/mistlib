// File: ./MetaverseClient.ts
import { WebRTCSignaling } from "./WebRTCSignaling";
import { WebSocketClient } from "./WebSocketClient";
import { MessageHandler } from "./MessageHandler";

const wsUrl = "ws://localhost:8080/ws"; // WebSocketサーバーのURLを指定
const wsClient = new WebSocketClient(wsUrl);
const messageHandler = new MessageHandler(wsClient);

const iceServers: RTCIceServer[] = [{ urls: "stun:stun.l.google.com:19302" }];
const webRTCSignaling = new WebRTCSignaling(wsUrl, iceServers);

// Example usage: WebRTCシグナリングの開始
webRTCSignaling.createOffer();

// Example usage: メッセージを送信する例
wsClient.sendMessage("RPC", { method: "someMethod", params: [1, 2, 3] });
