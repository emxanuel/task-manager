import { Component, Input, OnInit } from '@angular/core';
import { Task } from '../../core/models/task.interface';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';
import { TaskService } from '../../core/services/task.service';
import { EditTaskComponent } from '../edit-task/edit-task.component';

@Component({
  selector: 'task-item',
  imports: [HlmButtonDirective, EditTaskComponent],
  templateUrl: './task-item.component.html',
})
export class TaskItemComponent implements OnInit {
  @Input() task!: Task;
  date: string = '';

  constructor(private taskService: TaskService) {}

  ngOnInit(): void {
    if (this.task?.CreatedAt) {
      this.date = new Date(this.task.CreatedAt).toLocaleDateString();
    }
  }

  onDelete() {
    this.taskService.deleteTask(this.task.ID).subscribe({
      next: () => {
        this.taskService.dataStatus.emit('data updated');
      },
    });
  }
}
