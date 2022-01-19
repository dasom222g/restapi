if (window.EventSource) {
  const formArea = document.querySelector('#chat-form')
  const formNameArea = document.querySelector('#user-form')
  const nameArea = formArea.querySelector('.chat-name')
  const logArea = formArea.querySelector('.chat-log')
  
  const messageArea = formArea.querySelector('#chat-msg')
  const username = prompt('What is your name?')
  let userInfo = {}
  
  const handleSubmit = async (e) => {
    e.preventDefault()
    
    // 데이터 전송
    const message = messageArea.value
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
    messageArea.value = ''
    messageArea.focus()
  }
  const handleUserSubmit = async (e) => {
    e.preventDefault()

    try {
      const data = { name: username }
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
  es.onopen = async () => {
    try {
      const data = { name: username }
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

  es.onmessage = (e) => {
    // const data = e.data
    console.log('onmessage', e.data)
  }
  
  
  formArea.addEventListener('submit', handleSubmit)
  formNameArea.addEventListener('submit', handleUserSubmit)
}
