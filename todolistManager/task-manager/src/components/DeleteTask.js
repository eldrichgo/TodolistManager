import React, { useState } from 'react';
import axios from 'axios';

function DeleteTask() {
  const [taskID, setTaskID] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.delete(`http://localhost:8080/deletetask/${taskID}`);
      console.log('Task deleted');
    } catch (error) {
      console.error('Error deleting task:', error);
    }
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="text"
        value={taskID}
        onChange={(e) => setTaskID(e.target.value)}
        placeholder="Task ID"
      />
      <button type="submit">Delete Task</button>
    </form>
  );
}

export default DeleteTask;