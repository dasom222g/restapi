console.log(window.EventSource)

const formArea = document.querySelector('#chat-form')
const nameArea = formArea.querySelector('.chat-name')
const logArea = formArea.querySelector('.chat-log')

const messageArea = formArea.querySelector('#chat-msg')
const userName = prompt('What is your name?')

const handleSubmit = (e) => {
  e.preventDefault()
  const message = messageArea.value
  messageArea.value = ''
  messageArea.focus()

  console.log('message', message, 'userName', userName)
}


formArea.addEventListener('submit', handleSubmit)