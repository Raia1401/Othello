import React from "react"
import { useEffect,useContext,useState} from "react";
import { UserDataContext } from "../providers/UserDataProvider";
import { BoardDataContext} from "../providers/BoardDataProvider";
import * as apis from "../apis/api";
import type { BoardData } from "../types/BoardData";
import * as settings from "../settings/settings"


const convStringToArray = (rawStonesPos:string) => {
  let row=[]
  let col=[]
  for (let i=0; i< rawStonesPos.length; i++){
    if (rawStonesPos.charAt(i)!=="3"){
      let stoneColor = parseInt(rawStonesPos.charAt(i))
      row.push(stoneColor)
    }
    if ((i+1) % settings.BOARD_SIZE === 0 && row.length ){
      col.push(row)
      row=[]
    }
  }
  return col
}


export const useBoardData=()=>{

    // const userDataCtx=useContext(UserDataContext)
    const [boardId,setBoardId]=useState<number>(0)
    const [isMyTurn,setIsMyturn]=useState<boolean>(false)
    const [stonesPos,setStonesPos]=useState<number[][]>([])
    // const boardDataCtx=useContext(BoardDataContext)
    // userDataCtx.userId

    useEffect(()=>{
      apis.getBoardData(1).then((boardData:BoardData)=>{
        // console.log("boardData.data.Board",boardData.data.Board)
        // const stonePosArray = convStringToArray(boardData.data.Board)
        // console.log("stonePosArray",stonePosArray)
        setBoardId(boardData.data.BoardId)
        setStonesPos(convStringToArray(boardData.data.Board))
        setIsMyturn(boardData.data.IsMyTurn)
      })
    },[])
    // apis.getBoardData(userDataCtx.userId).then((boardData:BoardData)=>{
    //   const stonePosArray = convStringToArray(boardData.data.Board)
    //   console.log(stonePosArray)
    //   setStonesPos(stonePosArray)
    // })

    // useEffect(()=>{
    //   console.log(stonesPos)
    // },[stonesPos])
  
    
    // return boardDataCtx.stonesPos
    // return [boardId,isMyTurn,stonesPos];
    return {boardId:boardId,isMyTurn:isMyTurn,stonesPos:stonesPos}
}
