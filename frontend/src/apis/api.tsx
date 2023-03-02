import axios from "axios";
import type { StoneMovement } from "../types/StoneMovement";

const rootURL:string = process.env.REACT_APP_API || ""


//ユーザーのゲーム盤面を新規作成
export const postBoardData = async (userId:number)=>{
  // console.log(userId)
  const params = new URLSearchParams();
  params.append('user_id',userId.toString())
  const response =await axios.post(rootURL,params)
  return response.data
}

//ユーザーのゲーム盤面を取得
export const getBoardData = async (userId:number)=>{
  const response = await axios.get(`${rootURL}/${userId}`)
  return response.data
}

//ユーザーが石を置いて盤面を更新する処理
export const updateStonePos = async (boardId:number,stoneMovement:StoneMovement)=>{
  console.log(stoneMovement)
  const response = await axios.put(`${rootURL}/myself/${boardId}`,
    stoneMovement,
    {headers: {
      "Content-Type": "multipart/form-data",
    }})
  return response.data
}

//ユーザーの対戦相手（コンピューター）が石を置いて盤面を更新する処理
export const updateStonePosByOpponent = async (boardId:number)=>{
  const response = await axios.put(`${rootURL}/opponent/${boardId}`)
  return response.data
}