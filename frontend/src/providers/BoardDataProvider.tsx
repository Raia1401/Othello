import {createContext,useState} from "react"

const useStonesPos=()=>{
  const BOARD_SIZE=10
  const [stonesPos,_setStonesPos]=useState<number[][]>([])
  const convStringToArray = (rawStonesPos:string) => {
    let row=[]
    let col=[]
    for (let i=0; i< rawStonesPos.length; i++){
        let stoneColor = parseInt(rawStonesPos.charAt(i))
        row.push(stoneColor)
        if ((i+1) % BOARD_SIZE === 0){
            col.push(row)
            row=[]
        }
    }
    return col
  }
  const setStonesPos = (rawStonesPos:string) => _setStonesPos(convStringToArray(rawStonesPos))
  return {stonesPos,setStonesPos}
}

export const BoardDataContext = createContext({} as {
  boardId:number
  setBoardId:React.Dispatch<React.SetStateAction<number>>
  // stonesPos:number[][]
  // setStonesPos: (rawStonesPos: string) => void
  isMyTurn:boolean
  setIsMyTurn:React.Dispatch<React.SetStateAction<boolean>>
})

export const BoardDataProvider:React.FC<{children: React.ReactNode}> = (props)=>{

  const [boardId,setBoardId]=useState<number>(1)
  // const {stonesPos,setStonesPos}=useStonesPos()
  const [isMyTurn,setIsMyTurn]=useState<boolean>(true)

  // const [stonesPos,setStonesPos]=useState<number[][]>([])
  // const setStonesPos = (rawStonesPos:string) => _setStonesPos(convStringToArray(rawStonesPos))

  return (
    // <BoardDataContext.Provider value={{boardId,setBoardId,stonesPos,setStonesPos,isMyTurn,setIsMyTurn}}>
    <BoardDataContext.Provider value={{boardId,setBoardId,isMyTurn,setIsMyTurn}}>
        {props.children}
    </BoardDataContext.Provider>
  )
}