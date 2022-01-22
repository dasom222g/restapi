if (window.EventSource) {
  const nameFormArea = document.querySelector('.user-area')
  const nameForm = document.querySelector('#user-form')
  const messageFormArea = document.querySelector('.message-area')
  const messageForm = document.querySelector('#chat-form')
  const nameArea = messageForm.querySelector('.chat-name')
  const logArea = messageForm.querySelector('.chat-log')
  
  const messageInput = messageForm.querySelector('#chat-msg')
  const nameInput = nameForm.querySelector('#username')
  let userInfo = {}
  
  const handleMessageSubmit = async (e) => {
    e.preventDefault()
    
    // 데이터 전송
    const message = messageInput.value.trim()
    if (!message.length) return
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
    console.log('add user!!')
    const value = nameInput.value.trim()
    if (!value.length) return
    try {
      const data = { name: value }
      const response = await fetch('/user', {
        method: 'POST',
        body: JSON.stringify(data),
      })
      const result = await response.json()
      userInfo = result
      console.log('userInfo', userInfo)
      messageFormArea.style.display = 'block'
      nameFormArea.style.display = 'none'
      messageInput.focus()
      // console.log('userInfo', userInfo)
    } catch(error) {
      console.error(error)
    }
  }

  const addMessage = (data, type) => {
    const { id, name,  message } = data
    const element = document.createElement('div')
    element.innerHTML = message
    let align = 'left'
    if (type === 'userInfo') {
      // 입장 메시지
      element.style.backgroundColor = '#999'
      element.style.color = '#fff'
      align = 'center'
    }
    // 작성자가 본인일 경우
    align = Object.keys(userInfo).length && userInfo.id === id ? 'right' : align
    console.log('addMessage', userInfo)
    element.style.textAlign = align
    logArea.appendChild(element)
  }

  // eventSource
  const es = new EventSource('/stream')
  es.onopen = () => {
    init()
    console.log('eventSource open!!')
  }

  es.onmessage = (e) => {
    // const data = e.data
    const data = JSON.parse(e.data)
    console.log('on!!', data)
    addMessage(data, data.message.includes('입장') ? 'userInfo' : 'message')
  }
  
  messageForm.addEventListener('submit', handleMessageSubmit)
  nameForm.addEventListener('submit', handleUserInfoSubmit)

  const init = () => {
    messageFormArea.style.display = 'none'
    nameFormArea.style.display = 'block'
  }
}
