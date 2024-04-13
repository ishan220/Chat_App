import axios from 'axios'
import {useState} from 'react'
import {Navigate} from 'react-router-dom'
const LoginUser = () => {  
    const [credentials,setCredentials] = useState({})
    const [redirectParams,setRedirectParams] = useState({redirect:0,redirectTo:"/chat/?username="})
  
    const handleLogin = async(e) => {
        e.preventDefault();

        console.log("Inside handleLogin function")
        console.log("${process.env.REACT_APP_BACKEND_HOST}:",`${process.env.REACT_APP_BACKEND_HOST}/login`)
        await axios.post({headers:{
            'Access-Control-Allow-Origin': '*',
             'Accept': 'application/json',
            'Content-Type': 'application/json'
          },url:`${process.env.REACT_APP_BACKEND_HOST}/login`,data:{
            username:credentials.username,
            password:credentials.password
        }}).then((data)=>{
            console.log("data",data)
            const redirectTo = redirectParams.redirectTo + credentials.username
            setRedirectParams({redirect:1,redirectTo})
            }).catch((error)=>{
            console.log(" Unsuccessful login",error)  
            window.alert("UserName or Password is Wrong")
        });
        
    }
    
    console.log("rendering the comp",redirectParams.redirectTo,":",redirectParams.redirect)
  
    if(redirectParams && redirectParams.redirect) {
          return <><Navigate to={redirectParams.redirectTo} replace={true}/></>
        }
    return <> 
        <div>           
                <form className="register-form">
                 <input type="text" name="username" placeholder="Enter Username" value={(credentials && credentials.username) || ""} onChange={(e)=>{setCredentials({...credentials,"username":e.target.value})}}/>
                 <input type="password" name="password" placeholder="Enter Password!" value={(credentials && credentials.password) || ""} onChange={(e)=>{setCredentials({...credentials,"password":e.target.value})}}/>
                 <button onClick={(e)=>handleLogin(e)}>Login</button>
                </form>
          </div>
        </>
} 

export default LoginUser