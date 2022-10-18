import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { Router } from '@angular/router';
import { Topic } from 'src/app/models/topic.model';
import { ApiService } from 'src/app/services/api.service';
import { TopicSelectingService } from 'src/app/services/topic-selecting.service';
@Component({
  selector: 'app-topic-item',
  templateUrl: './topic-item.component.html',
  styleUrls: ['./topic-item.component.scss']
})
export class TopicItemComponent implements OnInit {
  @Input() topic?: Topic;
  @Output() selectedTopic = new EventEmitter<Topic>();

  constructor(private topicSelectingService: TopicSelectingService, private apiService: ApiService, private router: Router) { }

  onSelectTopic() {
    // if (this.topic) this.topicSelectingService.onGetTopic(this.topic);
    if (this.topic) {
      this.router.navigate(["/topic", this.topic.id])
    };
    // this.selectedTopic.emit(this.topic);
  }

  ngOnInit(): void {
  }

}
