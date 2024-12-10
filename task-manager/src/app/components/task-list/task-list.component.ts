import { Component, OnInit } from '@angular/core';
import { TaskItemComponent } from '../task-item/task-item.component';
import { GetTasksResponse } from '../../core/models/task.interface';
import { TaskService } from '../../core/services/task.service';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'task-list',
  imports: [TaskItemComponent, CommonModule],
  templateUrl: './task-list.component.html',
})

export class TaskListComponent implements OnInit {
  tasks: GetTasksResponse = [];
  loading: boolean = false;
  error: string = '';
  
  constructor(private taskService: TaskService) {}
  
  ngOnInit(): void {
    this.loading = true;
    
    this.taskService.getTasks().subscribe({
      next: (tasks: GetTasksResponse) => {
        this.tasks = tasks;
        this.loading = false;
      },
      error: () => {
        this.error = 'Failed to load tasks';
        this.loading = false;
      },
    });
  }
}
