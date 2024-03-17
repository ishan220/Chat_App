import { useParams,useLocation } from "react-router-dom"
import {useState, useEffect,useRef } from "react"
import {lazily} from 'react-lazily'
//import {connect,sendMsg,OnConnection,DisConnect,socket} from './socketconnection'
// import  from './socketconnection'
import Contacts  from "./Contact"
import ChatHistory  from "./ChatHistory"
import HandlingSocketConn from "./HandlingSocketConn.js"
import axios from 'axios'

const Chat = () => {
     const location = useLocation();
     const queryParams = new URLSearchParams(location.search)
     const user = queryParams.get('username')
     const chatRef = useRef();
       
     const [state,setState] = useState({username:user,
                                        renderedContactList:[],
                                        searchContact:"",
                                        contactList:[],
                                        to:"",
                                        renderedChatHistory:[]}); 

        const [chatHistory,setChatHistory]=useState([])        
        const [socketConn,setSocketConn]=useState() 
        chatRef.current=chatHistory
                 
        const [contactList,setContactList]  =useState([])                               
        const [message,setMessage] = useState("")


    const sendMessageTo = (to) => {
        console.log("On clicking on contact:",to)
        fetchChatHistory(state.username,to)
        setState((prevState)=>({...prevState,to}))
        }

    ///this is called on clicking on Add contact
    const addContact = async(e,searchContact) =>{
        e.preventDefault()
        await axios.post(`http://localhost:8080/addContact`,{
            username:state.username,
            last_activity:23121,
            member:state.searchContact
        }).then((response)=>{
            let contactsFromState= state.contactList
            contactsFromState.push(response.data)
            setState((prevState)=>({...prevState,contactList:contactsFromState}))
            renderContactList(contactsFromState)
            window.alert("user Added")
        }).catch((error)=>{
            console.log("User can't Be Added",error)
            window.alert(error)
        })
    }
    //this is called when in contact list any contact is clicked to
    const fetchChatHistory = async(from ,to) => {
      //  let newChatHistory=state.chatHistory;
        let newChatHistory =  chatRef.current;
          await axios.get(`http://localhost:8080/ChatBtwnUsers?from=${from}&to=${to}`)
         .then((response)=>{
            console.log("response data from fetch chat api",response.data)
            if(response.data){
            setChatHistory(response.data)
            renderChatHistory(response.data)
            }
            else
            {
            setChatHistory([])
            renderChatHistory([])
            }
            //let statsChatHistory=state.chatHistory.unshift()
         }).catch((error)=>{
            //console.log("Error from fetch chat Api",error)
         });
    }

     const getContacts = async(username)=>{
            let contacts=[];
            await axios.get(`http://localhost:8080/getContacts?username=${state.username}`)
            .then((response)=>{console.log("response",response.data);
                if(response.data && response.data.length>0)
                contacts=response.data
            } ).catch((error)=>{
           //     console.log("Error from get Contacts",error)
            }) 
           // setContactList({ContactList:contacts})
            setState((prevState)=>({...prevState,contactList:contacts}))
            console.log("Setted state.contactList")
            renderContactList(contacts);
        }

     const renderContactList =  (contacts) => {
        const renderContactList = Contacts(contacts,sendMessageTo)
        //console.log("renderContactList In renderContactList:",renderContactList)
         setState((prevState)=>({...prevState,renderedContactList:renderContactList}))
         //console.log("renderedContact list function End")
    }
   
    
    const OnSendChat = (e)=>{
        e.preventDefault()
       const msg ={type:"message",
                  user:state.username,
                  chat:{id:(Math.floor(Math.random() * 1000000)).toString(),
                        FromUser:user,
                        ToUser:state.to,
                        message:message,}
                     ,}                                  
         socketConn.send(JSON.stringify(msg))
         let currChatHist=[];
         if(chatHistory)
         currChatHist = chatHistory;
         currChatHist.push(msg.chat)
         setChatHistory(currChatHist)
         renderChatHistory(currChatHist)
        setMessage("");
    }
     useEffect(()=>{
              //console.log("Use effect getting called")
              const socket = new WebSocket("ws://localhost:8080/ws")
              setSocketConn(socket)
              HandlingSocketConn(socket,(msg)=>{
                                        console.log("Msg received from another client",msg.data) 
                                        const message = JSON.parse(msg.data);
                                        // console.log("message ka type",typeof(message))
                                        // console.log("msg ToUser: User",message.to_user,":",user)
                                        if(message.to_user == user && state.to == message.from_user)
                                        {
                                        console.log("fetch api called on incoming message")
                                        let currChatHist = chatRef.current;
                                        console.log("chat history before pushing",currChatHist)
                                        currChatHist.push(message)
                                        setChatHistory(currChatHist);
                                        renderChatHistory(currChatHist)
                                        }
                                 },()=>{
                                    const msg = {type:"Bootup",
                                    user:state.username,
                                    chat:{ FromUser:state.username,
                                    Message:message,
                                     },}    
                                    console.log("Bootup message being sent")
                                    socket.send(JSON.stringify(msg))
                                 })
             getContacts(state.username)
            },[])

        const renderChatHistory = (chatHistory) => {
            const renderedChatHistory = ChatHistory(chatHistory,user)
            setState((prev)=>({...prev,renderedChatHistory}))
        }

        console.log("chatHistory",chatHistory)

    return <>
    <div className="contact-search">
       
        <input type='text' placeholder='Search Contact' value={state.searchContact} 
                            onChange={(e)=>{ setState({...state,searchContact:e.target.value})}}>
        </input>

        <button style={{display:"inline-block"}} onClick={(e)=>{addContact(e,state.searchContact)}}>
             Add Contacts
        </button>

    </div>

    <div className="chat-container">
       
        <div className="contact-list">
            {state.renderedContactList && state.renderedContactList.length>0 && state.renderedContactList.map(contact=> contact)}
        </div>
      
        <div className="chat-area" >
                
                 <div className="chat-hist">
                    {state.renderedChatHistory}
                 </div>
            
                 <div className="chat-box">
                    <form>
                        <input type='text' onChange={(e)=>{setMessage(e.target.value)}} value={message} placeholder="Send Message"></input>
                        <button onClick={(e)=>(OnSendChat(e))}>Send</button>
                    </form>
                </div>
        </div>

    </div>

    </>
}
export default Chat


/////////
/*All call to setState is clubbed and executed once all
 setState is called unless there is axios requests for whichit will not wait */
//////////