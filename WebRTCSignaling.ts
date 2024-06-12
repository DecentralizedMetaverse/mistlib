// File: ./WebRTCSignaling.ts
import { WebSocketClient } from "./WebSocketClient";
import { WebRTCClient } from "./WebRTCClient";
import { MessageType } from "./MessageType";

export class WebRTCSignaling {
  private webSocketClient: WebSocketClient;
  private webRTCClient: WebRTCClient;

  constructor(wsUrl: string, iceServers: RTCIceServer[]) {
    this.webSocketClient = new WebSocketClient(wsUrl);
    this.webRTCClient = new WebRTCClient(iceServers, this.handleRTCMessage.bind(this), this.handleRTCConnectionStateChange.bind(this));
    this.registerMessageHandlers();
  }

  private registerMessageHandlers() {
    this.webSocketClient.addMessageHandler(MessageType.Offer, this.handleOffer.bind(this));
    this.webSocketClient.addMessageHandler(MessageType.Answer, this.handleAnswer.bind(this));
    this.webSocketClient.addMessageHandler(MessageType.Candidate, this.handleCandidate.bind(this));
  }

  public createOffer() {
    this.webRTCClient.createOffer();
    this.webSocketClient.sendMessage(MessageType.Offer, this.webRTCClient.peerConnection.localDescription);
  }

  public createAnswer(offer: RTCSessionDescriptionInit) {
    this.webRTCClient.createAnswer(offer);
    this.webSocketClient.sendMessage(MessageType.Answer, this.webRTCClient.peerConnection.localDescription);
  }

  private handleOffer(offer: any) {
    this.webRTCClient.createAnswer(offer);
    this.webSocketClient.sendMessage(MessageType.Answer, this.webRTCClient.peerConnection.localDescription);
  }

  private handleAnswer(answer: any) {
    this.webRTCClient.setRemoteDescription(answer);
  }

  private handleCandidate(candidate: any) {
    this.webRTCClient.addIceCandidate(candidate);
  }

  private handleRTCMessage(message: any) {
    console.log("Received RTC message:", message);
  }

  private handleRTCConnectionStateChange(state: RTCPeerConnectionState) {
    console.log("RTC connection state change:", state);
  }

  public sendMessage(type: string, data: any) {
    this.webRTCClient.sendMessage(type, data);
  }
}
