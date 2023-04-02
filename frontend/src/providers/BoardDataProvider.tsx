import {createContext,useState} from "react"

export const BoardDataContext = createContext({} as {
  boardId:number
  setBoardId:React.Dispatch<React.SetStateAction<number>>
  isMyTurn:boolean
  setIsMyTurn:React.Dispatch<React.SetStateAction<boolean>>
})

export const BoardDataProvider:React.FC<{children: React.ReactNode}> = (props)=>{

  const [boardId,setBoardId]=useState<number>(0)
  const [isMyTurn,setIsMyTurn]=useState<boolean>(true)

  return (
    <BoardDataContext.Provider value={{boardId,setBoardId,isMyTurn,setIsMyTurn}}>
        {props.children}
    </BoardDataContext.Provider>
  )
}