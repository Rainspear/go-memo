import { Component, EventEmitter, Input, OnInit, OnChanges, Output, SimpleChanges, OnDestroy } from '@angular/core';
import { faXmarkCircle } from '@fortawesome/free-solid-svg-icons';
import { Subject, Subscription } from 'rxjs';

@Component({
  selector: 'app-lightbox',
  templateUrl: './lightbox.component.html',
  styleUrls: ['./lightbox.component.scss']
})
export class LightboxComponent implements OnInit, OnChanges, OnDestroy {

  @Input() show?: boolean;
  @Output() toggle = new EventEmitter<boolean>();

  faXmarkCircle = faXmarkCircle;

  constructor() {
  }

  onClickCloseApp() {
    this.toggle.emit(!this.show);
  }

  ngOnChanges(changes: SimpleChanges): void {
    // let show = changes['show'].currentValue
    if (this.show) { document.body.style.overflow = 'hidden' }
    else {
      document.body.style.overflow = 'auto'
    }

  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
  }

}
