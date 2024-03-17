import '../css/style.css'  
const ChatHistory = (chatHistory,user)=> {

    
    const allChats =  chatHistory.map((chats,index) => {
    
        if( (chats.from_user && chats.from_user == user) 
            || (chats.FromUser && chats.FromUser==user) 
          )
             {
            return <div key={index} className="from-chat">{chats.message}</div>
             }
     
        else{
            return <div key={index} className="to-chat">{chats.message}</div>
        }

    })
    console.log("allChats",allChats)
    return  allChats
}

export default ChatHistory;