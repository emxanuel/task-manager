import { Component, Input, OnInit } from '@angular/core';
import { Task } from '../../core/models/task.interface';

@Component({
  selector: 'task-item',
  imports: [],
  templateUrl: './task-item.component.html',
})
export class TaskItemComponent implements OnInit {
  @Input() task!: Task
  date: string = ''

  constructor(){}

  ngOnInit(): void {
      if (this.task?.CreatedAt){
        this.date = new Date (this.task.CreatedAt).toLocaleDateString()
      }
  }
}
