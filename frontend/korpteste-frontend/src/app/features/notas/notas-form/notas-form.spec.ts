import { ComponentFixture, TestBed } from '@angular/core/testing';

import { NotasForm } from './notas-form';

describe('NotasForm', () => {
  let component: NotasForm;
  let fixture: ComponentFixture<NotasForm>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [NotasForm],
    }).compileComponents();

    fixture = TestBed.createComponent(NotasForm);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
