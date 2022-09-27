import * as jspb from 'google-protobuf'



export class BoyfriendRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BoyfriendRequest.AsObject;
  static toObject(includeInstance: boolean, msg: BoyfriendRequest): BoyfriendRequest.AsObject;
  static serializeBinaryToWriter(message: BoyfriendRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BoyfriendRequest;
  static deserializeBinaryFromReader(message: BoyfriendRequest, reader: jspb.BinaryReader): BoyfriendRequest;
}

export namespace BoyfriendRequest {
  export type AsObject = {
  }
}

export class BoyfriendResponse extends jspb.Message {
  getEmoji(): string;
  setEmoji(value: string): BoyfriendResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): BoyfriendResponse.AsObject;
  static toObject(includeInstance: boolean, msg: BoyfriendResponse): BoyfriendResponse.AsObject;
  static serializeBinaryToWriter(message: BoyfriendResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): BoyfriendResponse;
  static deserializeBinaryFromReader(message: BoyfriendResponse, reader: jspb.BinaryReader): BoyfriendResponse;
}

export namespace BoyfriendResponse {
  export type AsObject = {
    emoji: string,
  }
}

