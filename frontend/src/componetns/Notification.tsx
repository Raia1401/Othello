import "./Notification.css"


export const Notification:React.FC<{isMyTurn:boolean,children: React.ReactNode}>=(props)=>{

  // const BOARD_SIZE:number=8
  // const grid=[]
  // let row=[]

  // for (let x=0; x< BOARD_SIZE; x++){
  //   for (let y=0; y< BOARD_SIZE; y++){
  //     let stoneColor:number =props.stonesPos[x][y]
  //     let keyName:string = 'square:'+x.toString()+'_'+y.toString()
  //     row.push(<DrawSquare x={x} y={y} colorOfStone={stoneColor} key={keyName}/>)
  //   }
  //   let keyName:string = 'row:'+x.toString()
  //   grid.push(<div className='Grid-row' key={keyName}>{row}</div>)
  //   row=[]
  // }

  return (
    <div>
      <div className="content">
        {props.isMyTurn ?
        <div className="myturn-notification">
        <p>自分の番です</p>
        </div>
        :
        <div className="opponentturn-notification">
        <p>相手の番です</p>
        </div>
      }
      </div>
      {props.children}
    </div>
  )
}

export default Notification;
