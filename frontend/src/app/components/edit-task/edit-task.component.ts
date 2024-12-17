import { Component, Input, OnInit } from '@angular/core';
import {
  HlmSheetComponent,
  HlmSheetContentComponent,
  HlmSheetDescriptionDirective,
  HlmSheetTitleDirective,
  HlmSheetHeaderComponent
} from '@spartan-ng/ui-sheet-helm';
import { Task } from '../../core/models/task.interface';
import { HlmButtonDirective } from '@spartan-ng/ui-button-helm';

@Component({
  imports: [
    HlmSheetComponent,
    HlmSheetContentComponent,
    HlmSheetHeaderComponent,
    HlmSheetTitleDirective,
    HlmButtonDirective,
    HlmSheetDescriptionDirective,
  ],
  templateUrl: './edit-task.component.html',
  selector: 'edit-task',
})
export class EditTaskComponent implements OnInit {
  @Input() task!: Task;

  ngOnInit(): void {
      console.log(this.task)
  }
}
