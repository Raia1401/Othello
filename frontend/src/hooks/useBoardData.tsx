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

    const [stonesPos,setStonesPos]=useState<number[][]>([])
    const [isMatchEnd,setIsMatchEnd]=useState<boolean>(false)

    useEffect(()=>{
      apis.getBoardData(boardId).then((boardData:BoardData)=>{
        setStonesPos(convStringToArray(boardData.data.Board))
        setIsMatchEnd(boardData.data.IsMatchEnd)
      })
    },[boardId,isMyTurn])

    return {stonesPos:stonesPos,isMatchEnd:isMatchEnd}
}
