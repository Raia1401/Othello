import { useContext,useEffect } from "react"
import * as apis from "../apis/api"
import { BoardDataContext } from "../providers/BoardDataProvider"
import type { BoardData } from "../types/BoardData"

export const useOpponent=()=>{

    const boardDataCtx=useContext(BoardDataContext)
    
    useEffect(()=>{

        window.setTimeout(function(){
            console.log('タイマーが作動した')
        }, 5000);
        // window.setTimeout(function(){apis.updateStonePosByOpponent(boardDataCtx.boardId).then((boardData:BoardData)=>{
        //     boardDataCtx.setBoardId(boardData.data.BoardId)
        //     boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
        // })}, 5000);

        // apis.updateStonePosByOpponent(boardDataCtx.boardId).then((boardData:BoardData)=>{
        //     boardDataCtx.setBoardId(boardData.data.BoardId)
        //     boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
        // })
    },[])
    

    // const onClickBoard = (x:number,y:number,colorOfStone:number) =>{
    //     // console.log(x,y)
    //     if (true){
    //         const stoneMovement={"x":x,"y":y,"color":colorOfStone}
    //         apis.updateStonePos(boardDataCtx.boardId, stoneMovement).then((boardData:BoardData)=>{
    //             boardDataCtx.setBoardId(boardData.data.BoardId)
    //             boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
    //         })
    //     }
    // }
}