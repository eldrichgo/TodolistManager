import React from 'react';
import AddTask from './components/AddTask';
import ViewTasks from './components/ViewTasks';
import UpdateTask from './components/UpdateTask';
import DeleteTask from './components/DeleteTask';

function App() {
  return (
    <div>
      <h1>Task Manager</h1>
      <AddTask />
      <ViewTasks />
      <UpdateTask />
      <DeleteTask />
    </div>
  );
}

export default App;