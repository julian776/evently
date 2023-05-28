import { ComponentFixture, TestBed } from '@angular/core/testing';

import { IndividualEventComponent } from './individual-event.component';

describe('IndividualEventComponent', () => {
  let component: IndividualEventComponent;
  let fixture: ComponentFixture<IndividualEventComponent>;

  beforeEach(() => {
    TestBed.configureTestingModule({
      declarations: [IndividualEventComponent]
    });
    fixture = TestBed.createComponent(IndividualEventComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
