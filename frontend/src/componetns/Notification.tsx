import "./Notification.css"
import {TailSpin} from 'react-loader-spinner'


export const Notification:React.FC<{isMyTurn:boolean,children: React.ReactNode}>=(props)=>{

  return (
    <div>
      <div className="content">
        {props.isMyTurn ?
        <div className="myturn-notification">
        <p>自分の番です</p>
        </div>
        :
        <div className="opponentturn-notification">
          相手の番です
        <TailSpin width={30} color="#4fa94d" ariaLabel="tail-spin-loading" radius="1" visible={true} />
        </div>
      }
      </div>
      {props.children}
    </div>
  )
}

export default Notification;
