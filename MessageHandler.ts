// File: ./MessageHandler.ts
import { WebSocketClient } from "./WebSocketClient";
import { MessageType } from "./MessageType";

export class MessageHandler {
  private client: WebSocketClient;

  constructor(client: WebSocketClient) {
    this.client = client;
    this.registerHandlers();
  }

  private registerHandlers() {
    this.client.addMessageHandler(MessageType.RPC, this.handleRPC.bind(this));
    this.client.addMessageHandler(MessageType.Location, this.handleLocation.bind(this));
    this.client.addMessageHandler(MessageType.Animation, this.handleAnimation.bind(this));
    this.client.addMessageHandler(MessageType.Message, this.handleMessage.bind(this));
    this.client.addMessageHandler(MessageType.ObjectInstantiate, this.handleObjectInstantiate.bind(this));
    this.client.addMessageHandler(MessageType.Signaling, this.handleSignaling.bind(this));
    this.client.addMessageHandler(MessageType.PeerData, this.handlePeerData.bind(this));
    this.client.addMessageHandler(MessageType.Ping, this.handlePing.bind(this));
    this.client.addMessageHandler(MessageType.Pong, this.handlePong.bind(this));
  }

  private handleRPC(data: any) {
    console.log("Handle RPC", data);
    // RPC処理の実装
  }

  private handleLocation(data: any) {
    console.log("Handle Location", data);
    // Location処理の実装
  }

  private handleAnimation(data: any) {
    console.log("Handle Animation", data);
    // Animation処理の実装
  }

  private handleMessage(data: any) {
    console.log("Handle Message", data);
    // Message処理の実装
  }

  private handleObjectInstantiate(data: any) {
    console.log("Handle Object Instantiate", data);
    // Object Instantiate処理の実装
  }

  private handleSignaling(data: any) {
    console.log("Handle Signaling", data);
    // Signaling処理の実装
  }

  private handlePeerData(data: any) {
    console.log("Handle Peer Data", data);
    // Peer Data処理の実装
  }

  private handlePing(data: any) {
    console.log("Handle Ping", data);
    // Ping処理の実装
  }

  private handlePong(data: any) {
    console.log("Handle Pong", data);
    // Pong処理の実装
  }
}
