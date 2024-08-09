import React, { useState, useEffect } from 'react';

const ToDo = () => {
    const [todos, setTodos] = useState([]);
    const [task, setTask] = useState('');
    // const apiUrl = process.env.REACT_APP_API_URL;
    const apiUrl = "http://localhost:8000";

    useEffect(() => {
        fetch(`${apiUrl}/todos`)
            .then(response => response.json())
            .then(data => setTodos(data))
            .catch(error => console.error('Error fetching todos:', error));
    }, [apiUrl]);

    const addToDo = () => {
        const newToDo = { id: Date.now().toString(), task, done: false };
        console.log('adding todo: ', newToDo);
        fetch(`${apiUrl}/todos`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(newToDo),
        })
            .then(response => response.json())
            .then(data => setTodos([...todos, data]))
            .catch(error => console.error('Error adding todo:', error));
        setTask('');
    };

    const updateToDo = id => {
        const updatedToDo = todos.find(todo => todo.id === id);
        updatedToDo.done = !updatedToDo.done;
        fetch(`${apiUrl}/todos/${id}`, {
            method: 'PUT',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(updatedToDo),
        })
            .then(response => response.json())
            .then(() => {
                const updatedToDos = todos.map(todo =>
                    todo.id === id ? updatedToDo : todo
                );
                setTodos(updatedToDos);
            })
            .catch(error => console.error('Error updating todo:', error));
    };

    const deleteToDo = id => {
        console.log('deleting todo: ', id);
        fetch(`${apiUrl}/todos/${id}`, { method: 'DELETE' })
            .then(() => {
                const remainingToDos = todos.filter(todo => todo.id !== id);
                setTodos(remainingToDos);
            })
            .catch(error => console.error('Error deleting todo:', error));
    };

    return (
        <div>
            <h1>ToDo List</h1>
            <input
                type="text"
                value={task}
                onChange={e => setTask(e.target.value)}
            />
            <button onClick={addToDo}>Add Task</button>
            <button onClick={deleteToDo}>Delete Task</button>
            <ul>
                {todos.map(todo => (
                    <li key={todo.id}>
                        <span
                            style={{
                                textDecoration: todo.done ? 'line-through' : 'none',
                            }}
                        >
                            {todo.task}
                        </span>
                        <button onClick={() => updateToDo(todo.id)}>
                            {todo.done ? 'Undo' : 'Done'}
                        </button>
                        <button onClick={() => deleteToDo(todo.id)}>Delete</button>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ToDo;
