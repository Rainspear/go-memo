import { EventEmitter, Component, OnInit, Input, Output } from '@angular/core';
import { CreateParamsTopic, Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';

@Component({
  selector: 'app-topic-creating-modal',
  templateUrl: './topic-creating-modal.component.html',
  styleUrls: ['./topic-creating-modal.component.scss']
})
export class TopicCreatingModalComponent implements OnInit {

  constructor(private apiService: ApiService) { }

  @Input() show: boolean = false;
  @Output() toggle = new EventEmitter<boolean>()
  topic: CreateParamsTopic = { title: '', description: '' }
  error?: string;

  onToggleHandler(show: boolean) {
    this.toggle.emit(show);
  }

  onSubmitHandler() {
    this.error = undefined;
    if (this.topic.title && this.topic.description) {
      this.apiService.createTopic(this.topic).subscribe((res: any) => {
        if (res.data) {
          // console.log("res", res.data)
          this.toggle.emit(false)
        }
      }, (error) => {
        this.error = error.error
      })
    }
  }

  ngOnInit(): void {
  }



}
