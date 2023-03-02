import { useContext } from "react"
import * as apis from "../apis/api"
import { UserDataContext } from "../providers/UserDataProvider"
import { BoardDataContext } from "../providers/BoardDataProvider"
import type { BoardData } from "../types/BoardData"

export const useStoneMovement=()=>{

    const userDataCtx=useContext(UserDataContext)
    const boardDataCtx=useContext(BoardDataContext)

    const onClickBoard = (x:number,y:number,colorOfStone:number) =>{
        // console.log(x,y)
        if (boardDataCtx.isMyTurn){
            const stoneMovement={"user_id":userDataCtx.userId,"x":x,"y":y,"color":colorOfStone}
            apis.updateStonePos(boardDataCtx.boardId, stoneMovement).then((boardData:BoardData)=>{
                window.location.reload()
            //     boardDataCtx.setStonesPos(boardData.data.Board)
            })
        }
    }

    return onClickBoard

}