import React, { useEffect, useState } from 'react';
import axios from 'axios';

function ViewTasks() {
  const [tasks, setTasks] = useState([]);

  useEffect(() => {
    const fetchTasks = async () => {
      try {
        const response = await axios.get('http://localhost:8080/viewtasks');
        setTasks(response.data);
      } catch (error) {
        console.error('Error fetching tasks:', error);
      }
    };

    fetchTasks();
  }, []);

  return (
    <div>
      <h2>Tasks</h2>
      <ul>
        {tasks.map((task) => (
          <li key={task.ID}>
            ID: {task.ID} | Title: {task.Title} | Status: {task.Status}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default ViewTasks;
