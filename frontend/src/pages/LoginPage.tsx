import Grid from "../componetns/Grid";
import { useBoardData } from "../hooks/useBoardData";
import {Loding} from "../componetns/Loding"
import { postBoardData } from "../apis/api";
import './LoginPage.css'

function LoginPage(){

  const onClick=()=>{
    postBoardData(1)
    window.location.href = "/game"
  }

  return (
    <div>
      <h1>オセロゲーム</h1>
      <div className="button" onClick={() => onClick()}>新規ゲームを始める</div>
    </div>

  )

}


export default LoginPage;
