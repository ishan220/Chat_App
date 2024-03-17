// import {useNavigate} from 'react-router-dom'
import {Link} from 'react-router-dom'
import  '../css/style.css'


const Landing = () =>{
    return <> 
    <div className="landing-content">
            <div className="register">
                <Link to ="./register">Register</Link>
                {/* <button onClick = {()=>{handleRegister()}}>Register</button> */}
            </div>
            <div className="login">
            <Link to ="./login">Login</Link>
            </div>
    </div> 
    </>
}
export default Landing