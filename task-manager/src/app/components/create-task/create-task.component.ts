import { Component, ElementRef, ViewChild } from '@angular/core';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';
import { HlmInputDirective } from '@spartan-ng/ui-input-helm';
import { TaskService } from '../../core/services/task.service';
import { CreateTaskPayload } from '../../core/models/task.interface';

@Component({
  selector: 'create-task',
  imports: [HlmInputDirective, HlmButtonDirective],
  templateUrl: './create-task.component.html',
})
export class CreateTaskComponent {
  @ViewChild('taskTitleInput', { static: false })
  taskTitleInput!: ElementRef<HTMLInputElement>;
  @ViewChild('taskDescriptionInput', { static: false })
  taskDescriptionInput!: ElementRef<HTMLTextAreaElement>;

  isLoading: boolean = false;
  isSuccess: boolean = false;
  isError: boolean = false;
  error: string = '';

  constructor(private taskService: TaskService) {}

  onSubmit(event: Event) {
    const taskTitle = this.taskTitleInput.nativeElement.value.trim();
    const taskDescription =
      this.taskDescriptionInput.nativeElement.value.trim();
    event.preventDefault();
    if (!taskTitle) {
      this.isError = true;
      this.error = "Task's title can not be empty";
      return;
    }

    const newTask: CreateTaskPayload = {
      title: taskTitle,
      description: taskDescription,
    };

    this.isLoading = true;
    this.taskService.addTask(newTask).subscribe({
      next: () => {
        this.taskTitleInput.nativeElement.value = ''
        this.taskDescriptionInput.nativeElement.value = ''
        
        this.isLoading = false;
      },
      error: () => {
        this.taskTitleInput.nativeElement.value = ''
        this.taskDescriptionInput.nativeElement.value = ''

        this.isLoading = false;
        this.isError = true;
      },
    });
  }
}
