import {createContext,useState} from "react"

export const BoardDataContext = createContext({} as {
  boardId:number
  setBoardId:React.Dispatch<React.SetStateAction<number>>
  // stonesPos:number[][]
  // setStonesPos: (rawStonesPos: string) => void
  isMyTurn:boolean
  setIsMyTurn:React.Dispatch<React.SetStateAction<boolean>>
})

export const BoardDataProvider:React.FC<{children: React.ReactNode}> = (props)=>{

  const [boardId,setBoardId]=useState<number>(0)
  // const {stonesPos,setStonesPos}=useStonesPos()
  const [isMyTurn,setIsMyTurn]=useState<boolean>(true)

  // const [stonesPos,setStonesPos]=useState<number[][]>([])
  // const setStonesPos = (rawStonesPos:string) => _setStonesPos(convStringToArray(rawStonesPos))

  return (
    // <BoardDataContext.Provider value={{boardId,setBoardId,stonesPos,setStonesPos,isMyTurn,setIsMyTurn}}>
    // <BoardDataContext.Provider value={{boardId,setBoardId}}>
    <BoardDataContext.Provider value={{boardId,setBoardId,isMyTurn,setIsMyTurn}}>
        {props.children}
    </BoardDataContext.Provider>
  )
}