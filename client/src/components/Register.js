import axios from 'axios'
import {useState} from 'react'
import {Navigate} from 'react-router-dom'

const Register = () => {
    const [credentials,setCredentials] = useState({})
    const [redirectParams,setRedirectParams] = useState({redirect:false,redirectTo:"/login"})

    const handleRegistration = async(e) => {
        e.preventDefault()
        console.log("Inside handleRegistration function")
        await axios.post(`${process.env.BACKEND_HOST}/create-user`,{
            username:credentials.username,
            password:credentials.password
        }).then((resolved)=>{
            console.log("Registration Successful")
            window.alert("User Registered Successfully")
            console.log("redirectParams.redirect",redirectParams.redirect)
            console.log("redirectParams.redirectTo",redirectParams.redirectTo)
            const redirectTo=redirectParams.redirectTo
            setRedirectParams({redirect:true,redirectTo})

        }).catch((error)=>{
            window.alert("User Already Registered")
            console.log("Registration Unsuccessful",error)  
        });
    }

    if(redirectParams.redirect) 
    return  <><Navigate to={redirectParams.redirectTo} replace={true}/></>
    
    return <div>
        <form className="register-form">
        <input type="text" name="username" placeholder="Enter Username" value={(credentials && credentials.username) || ""} onChange={(e)=>{setCredentials({...credentials,"username":e.target.value})}}/>
        <input type="password" name="password" placeholder="Enter Password!" value={(credentials && credentials.password) || ""} onChange={(e)=>{setCredentials({...credentials,"password":e.target.value})}}/>
        <button onClick={(e)=>handleRegistration(e)}>Register</button>
        </form>
    </div>
}
export default Register;