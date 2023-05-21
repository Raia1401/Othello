import { useContext } from "react"

import * as apis from "../apis/api"
import { BoardDataContext } from "../providers/boardDataProvider"
import type { BoardData } from "../types/boardData"
import * as settings from "../settings/settings"

export const useTurnPass=()=>{

    const boardDataCtx=useContext(BoardDataContext)

    const passMyTurn = () =>{

        const stoneMovement={"x":0,"y":0,"my_stone_color":settings.MY_STONE_COLOR}
        apis.updateStonePos(boardDataCtx.boardId, stoneMovement).then((boardData:BoardData)=>{
            boardDataCtx.setBoardId(boardData.data.BoardId)
            boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
        })

        window.setTimeout(function(){
            apis.updateStonePosByOpponent(boardDataCtx.boardId,settings.MY_STONE_COLOR).then((boardData:BoardData)=>{
                    boardDataCtx.setBoardId(boardData.data.BoardId)
                    boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
                })
        }, 1000)

    }

    return passMyTurn

}