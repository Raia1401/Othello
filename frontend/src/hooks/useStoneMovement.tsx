import { useContext } from "react"

import * as apis from "../apis/api"
import { BoardDataContext } from "../providers/boardDataProvider"
import type { BoardData } from "../types/boardData"

export const useStoneMovement=()=>{

    const boardDataCtx=useContext(BoardDataContext)

    const onClickBoard = (x:number,y:number,colorOfStone:number) =>{

        const stoneMovement={"x":x,"y":y,"my_stone_color":colorOfStone}
        apis.updateStonePos(boardDataCtx.boardId, stoneMovement).then((boardData:BoardData)=>{
            boardDataCtx.setBoardId(boardData.data.BoardId)
            boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
            console.log("updateStonePos",boardData.data.IsMyTurn)
        })

        window.setTimeout(function(){
            apis.updateStonePosByOpponent(boardDataCtx.boardId,colorOfStone).then((boardData:BoardData)=>{
                    boardDataCtx.setBoardId(boardData.data.BoardId)
                    boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
                })
        }, 1000)

    }

    return onClickBoard

}