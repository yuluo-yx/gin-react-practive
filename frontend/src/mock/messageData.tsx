interface IMsg {
  id: number;
  user_id: number;
  user_name: string;
  text: string;
  item:Object[];
}

export default IMsg;
