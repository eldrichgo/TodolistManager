import React, { useState } from 'react';
import axios from 'axios';

function UpdateTask() {
  const [taskID, setTaskID] = useState('');
  const [status, setStatus] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.patch(`http://localhost:8080/updatetask/${taskID}`, { status });
      console.log('Task updated:', response.data);
    } catch (error) {
      console.error('Error updating task:', error);
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
      <input
        type="text"
        value={status}
        onChange={(e) => setStatus(e.target.value)}
        placeholder="New Status"
      />
      <button type="submit">Update Task</button>
    </form>
  );
}

export default UpdateTask;