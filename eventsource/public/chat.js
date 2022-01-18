if (window.EventSource) {
  const formArea = document.querySelector('#chat-form')
  const nameArea = formArea.querySelector('.chat-name')
  const logArea = formArea.querySelector('.chat-log')
  
  const messageArea = formArea.querySelector('#chat-msg')
  const username = prompt('What is your name?')
  
  const handleSubmit = (e) => {
    e.preventDefault()
    const message = messageArea.value
    messageArea.value = ''
    messageArea.focus()
  
    console.log('message', message, 'username', username)
  }
  
  const es = new EventSource('/stream')
  es.onopen = async () => {
    try {
      const data = { name: username }
      const response = await fetch('/user', {
        method: 'POST',
        body: JSON.stringify(data),
      })
      const result = await response.json()
      console.log('result', result)
    } catch(error) {
      console.error(error)
    }
  }
  
  
  formArea.addEventListener('submit', handleSubmit)
}
