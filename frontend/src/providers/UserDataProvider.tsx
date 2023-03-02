import {createContext,useState} from "react"
import * as settings from "../settings/settings"


export const UserDataContext = createContext({} as {
  userId:number,
  setUserId:React.Dispatch<React.SetStateAction<number>>
})

export const UserDataProvider:React.FC<{children: React.ReactNode}> = (props)=>{

  const [userId,setUserId]=useState<number>(settings.defaultUserId)

  return (
    <UserDataContext.Provider value={{userId,setUserId}}>
      {props.children}
    </UserDataContext.Provider>
  )
}