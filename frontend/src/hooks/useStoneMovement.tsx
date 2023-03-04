import { useContext } from "react"
import * as apis from "../apis/api"
import { BoardDataContext } from "../providers/BoardDataProvider"
import type { BoardData } from "../types/BoardData"

export const useStoneMovement=()=>{

    const boardDataCtx=useContext(BoardDataContext)

    const onClickBoard = (x:number,y:number,colorOfStone:number) =>{

        //TODO：自分のターンかどうかを把握する処理
        if (true){
            const stoneMovement={"x":x,"y":y,"color":colorOfStone}
            apis.updateStonePos(boardDataCtx.boardId, stoneMovement).then((boardData:BoardData)=>{
                boardDataCtx.setBoardId(boardData.data.BoardId)
                boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
                console.log("updateStonePos",boardData.data.IsMyTurn)
            })
        }

        window.setTimeout(function(){
            console.log('タイマーが作動した')
            apis.updateStonePosByOpponent(boardDataCtx.boardId,colorOfStone).then((boardData:BoardData)=>{
                    boardDataCtx.setBoardId(boardData.data.BoardId)
                    boardDataCtx.setIsMyTurn(boardData.data.IsMyTurn)
                })
        }, 1000)

    }

    return onClickBoard

}