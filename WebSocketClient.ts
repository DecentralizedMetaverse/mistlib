// File: ./WebSocketClient.ts
export class WebSocketClient {
    private ws: WebSocket;
    private url: string;
    private messageHandlers: { [key: string]: Function } = {};
    private reconnectInterval: number = 5000; // 再接続の間隔（ミリ秒）
  
    constructor(url: string) {
      this.url = url;
      this.connect();
    }
  
    private connect() {
      this.ws = new WebSocket(this.url);
      this.ws.onmessage = this.onMessage.bind(this);
      this.ws.onopen = () => console.log("Connected to WebSocket server");
      this.ws.onclose = () => {
        console.log("Disconnected from WebSocket server");
        setTimeout(() => this.connect(), this.reconnectInterval);
      };
    }
  
    private onMessage(event: MessageEvent) {
      const message = JSON.parse(event.data);
      const { type, data } = message;
      if (this.messageHandlers[type]) {
        this.messageHandlers[type](data);
      } else {
        console.log(`No handler for message type: ${type}`);
      }
    }
  
    public sendMessage(type: string, data: any) {
      this.ws.send(JSON.stringify({ type, data }));
    }
  
    public addMessageHandler(type: string, handler: Function) {
      this.messageHandlers[type] = handler;
    }
  }
  