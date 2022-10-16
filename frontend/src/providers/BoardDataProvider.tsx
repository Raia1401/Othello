import {createContext,useState} from "react"

const useStonesPos=()=>{
  const BOARD_SIZE=5
  const stonesPos:number[][]=[]
  const setStonesPos =(rawStonesPos:string)=>{
      let row=[]
      for (let i=0; i< rawStonesPos.length; i++){
          let stoneColor = parseInt(rawStonesPos.charAt(i))
          row.push(stoneColor)
          if (i %= BOARD_SIZE){
              stonesPos.push(row)
              row=[]
          }
      }
  }
  return {stonesPos,setStonesPos}
}

export const BoardDataContext = createContext({} as {
  boardId:number
  setBoardId:React.Dispatch<React.SetStateAction<number>>
  stonesPos:number[][]
  setStonesPos: (rawStonesPos: string) => void
  isMyTurn:boolean
  setIsMyTurn:React.Dispatch<React.SetStateAction<boolean>>
})

export const BoardDataProvider:React.FC<{children: React.ReactNode}> = (props)=>{

  const [boardId,setBoardId]=useState<number>(1)
  const {stonesPos,setStonesPos}=useStonesPos()
  const [isMyTurn,setIsMyTurn]=useState<boolean>(true)

  return (
    <BoardDataContext.Provider value={{boardId,setBoardId,stonesPos,setStonesPos,isMyTurn,setIsMyTurn}}>
      {props.children}
    </BoardDataContext.Provider>
  )
}