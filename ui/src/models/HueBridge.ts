class HueBridge {
  name: string;
  ipAddress: string;
  id: string;

  constructor(data: any) {
    this.id = data.id;
    this.name = data.name;
    this.ipAddress = data.ipAddress;
  }
}

export default HueBridge;
