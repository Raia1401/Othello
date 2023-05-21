import { useEffect, useState } from "react";

import * as apis from "../apis/api";
import * as settings from "../settings/settings";
import type { BoardData } from "../types/boardData";


const convStringToArray = (rawStonesPos:string) => {
  let row=[]
  let col=[]
  for (let i=0; i< rawStonesPos.length; i++){
    if (rawStonesPos.charAt(i)!=="3"){
      let stoneColor = parseInt(rawStonesPos.charAt(i))
      row.push(stoneColor)
    }
    if ((i+1) % (settings.BOARD_SIZE+2) === 0 && row.length ){
      col.push(row)
      row=[]
    }
  }
  return col
}


export const useBoardData=(boardId:number,isMyTurn:boolean)=>{

    // const userDataCtx=useContext(UserDataContext)
    // const [boardId,setBoardId]=useState<number>(0)
    // const [isMyTurn,setIsMyturn]=useState<boolean>(false)
    const [stonesPos,setStonesPos]=useState<number[][]>([])
    const [isMatchEnd,setIsMatchEnd]=useState<boolean>(false)

    // const {boardId,setBoardId} = useContext(BoardDataContext)
    // userDataCtx.userId

    useEffect(()=>{
      apis.getBoardData(boardId).then((boardData:BoardData)=>{
        // console.log("boardData.data.Board",boardData.data.Board)
        // const stonePosArray = convStringToArray(boardData.data.Board)
        // console.log("stonePosArray",stonePosArray)
        // setBoardId(boardData.data.BoardId)
        setStonesPos(convStringToArray(boardData.data.Board))
        setIsMatchEnd(boardData.data.IsMatchEnd)
        // setIsMyturn(boardData.data.IsMyTurn)
      })
    },[boardId,isMyTurn])

    return {stonesPos:stonesPos,isMatchEnd:isMatchEnd}
}
