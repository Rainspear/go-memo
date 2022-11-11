import { Component, EventEmitter, Input, OnInit, Output, SimpleChanges } from '@angular/core';
import { Topic } from 'src/app/models/topic.model';

@Component({
  selector: 'app-topic-list',
  templateUrl: './topic-list.component.html',
  styleUrls: ['./topic-list.component.scss']
})
export class TopicListComponent implements OnInit {
  showCreatingForm: boolean = false;
  test :number = 0
  @Input() topics: Topic[] = [];
  @Output() selectedTopic = new EventEmitter<Topic>();

  onSelectTopic(topic: Topic): void {
    this.selectedTopic.emit(topic);
  }

  onToggleForm() {
    console.log(this.showCreatingForm, "before")
    this.showCreatingForm = !this.showCreatingForm; 
    console.log(this.showCreatingForm, "after")
  }


  ngOnInit(): void {
  }

  ngOnChanges(changes: SimpleChanges) {
    // console.log("changes", changes)
    // console.log("TopicListComponent ngOnChanges: ", changes)
    // this.topics = changes['topics']
  }

}
