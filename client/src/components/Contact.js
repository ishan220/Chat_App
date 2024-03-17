// const contactEle =document.getElementsByClassName("contact")
// contactEle.addEventListener('onclick',function(){

// })
    // specify the action to take when the div is clicked
 import '../css/style.css'  
const Contacts = (contacts,sendMsgTo) =>  {
    const contactList = contacts.map((contact,index)=>{
        return (<div key={index} className="contact" 
                onClick={()=> {sendMsgTo(contact.member_username) }}>
                {contact.member_username} 
                </div>)
    })
    return contactList
}

export default Contacts