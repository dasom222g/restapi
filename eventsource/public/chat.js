if (window.EventSource) {
  const messageForm = document.querySelector('#chat-form')
  const nameForm = document.querySelector('#user-form')
  const nameArea = messageForm.querySelector('.chat-name')
  const logArea = messageForm.querySelector('.chat-log')
  
  const messageInput = messageForm.querySelector('#chat-msg')
  const nameInput = nameForm.querySelector('#username')
  let userInfo = {}
  
  const handleMessageSubmit = async (e) => {
    e.preventDefault()
    
    // 데이터 전송
    const message = messageInput.value
    const {id, name} = userInfo
    const messageInfo = {
      id,
      name,
      message
    }
    try {
      const response = await fetch('/message', {
        method: 'POST',
        body: JSON.stringify(messageInfo),
      })
      const result = await response.json()
      console.log('result', result)
    } catch(error) {
      console.error(error)
    }

    // ui 초기화
    messageInput.value = ''
    messageInput.focus()
  }
  const handleUserInfoSubmit = async (e) => {
    e.preventDefault()
    const value = nameInput.value.trim()
    if (!value.length) return
    try {
      const data = { name: value }
      const response = await fetch('/user', {
        method: 'POST',
        body: JSON.stringify(data),
      })
      const result = await response.json()
      console.log('result', result)
      userInfo = result
    } catch(error) {
      console.error(error)
    }
  }
  
  const es = new EventSource('/stream')
  es.onopen = () => {
    console.log('eventSource open!!')
  }

  es.onmessage = (e) => {
    // const data = e.data
    console.log('onmessage', e.data)
  }
  
  
  messageForm.addEventListener('submit', handleMessageSubmit)
  nameForm.addEventListener('submit', handleUserInfoSubmit)
}
