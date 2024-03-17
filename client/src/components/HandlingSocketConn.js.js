const HandlingSocketConn = (socket,cb1,cb2)=>{
    console.log("socket created:",socket)

    socket.onopen=()=>{
        console.log("Socket Connected");
        cb2()
    }

    socket.onmessage=(msg)=>{
        console.log("Message Received");
        cb1(msg)
    }

    socket.onclose=()=>{
        console.log("Connection Closed")
    }

    socket.onerror=(error)=>{
        console.log("Socket error",error)
    }
}
export default HandlingSocketConn