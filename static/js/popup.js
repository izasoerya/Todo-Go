function buttonCreate() {
  alert("Todo Created!")
}
function buttonDelete() {
  alert("Todo Deleted!")
}
function buttonEdit() {
  alert("Todo Edited!")
}
function buttonShow() {
  alert("Todo Showed!")
}
function handleKeyDown(event) {
  if (event.keyCode === 13 || event.key === 'Enter') {
      alert("Todo Created!")
  }
}
function readMessage() {
  fetch('./api/Todos/createTodo')
      .then((response) => response.json())
  alert(response)
}


  
