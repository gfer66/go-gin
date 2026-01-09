function changeMessage(){
   const messageElement = document.getElementById("message");
   const currentMessage = messageElement.textContent;

   const alternateMessages = [
      "Hello again!",
      "Welcome back!",
      "Have a nice day!"
   ];

   let newMessage;

   do {
      newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)]
   } while (newMessage === currentMessage);

   messageElement.textContent = newMessage;
}