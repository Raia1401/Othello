import React from "react"
import { useEffect,useContext } from "react";
import { UserDataContext } from "../providers/UserDataProvider";
import { BoardDataContext } from "../providers/BoardDataProvider";
import * as apis from "../apis/api";
import type { BoardData } from "../types/BoardData";


export const useStonesPos=()=>{

    const userDataCtx=useContext(UserDataContext)
    const boardDataCtx=useContext(BoardDataContext)

    useEffect(()=>{
        apis.getBoardData(userDataCtx.userId).then((boardData:BoardData)=>{
          console.log(boardData.board)
          boardDataCtx.setStonesPos(boardData.board)
        })
      })
    
    return boardDataCtx.stonesPos
}


// export const useStonesPos=()=>{
//     const BOARD_SIZE=5
//     const stonesPos:number[][]=[]
//     const setStonesPos =(rawStonesPos:string)=>{
//         let row=[]
//         for (let i=0; i< rawStonesPos.length; i++){
//             let stoneColor = parseInt(rawStonesPos.charAt(i))
//             row.push(stoneColor)
//             if (i %= BOARD_SIZE){
//                 stonesPos.push(row)
//                 row=[]
//             }
//         }
//     }
//     return {stonesPos,setStonesPos}
// }