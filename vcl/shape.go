
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
    "unsafe"
)

type TShape struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewShape
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewShape(owner IComponent) *TShape {
    s := new(TShape)
    s.instance = Shape_Create(CheckPtr(owner))
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ShapeFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ShapeFromInst(inst uintptr) *TShape {
    s := new(TShape)
    s.instance = inst
    s.ptr = unsafe.Pointer(inst)
    return s
}

// ShapeFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ShapeFromObj(obj IObject) *TShape {
    s := new(TShape)
    s.instance = CheckPtr(obj)
    s.ptr = unsafe.Pointer(s.instance)
    return s
}

// ShapeFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ShapeFromUnsafePointer(ptr unsafe.Pointer) *TShape {
    s := new(TShape)
    s.instance = uintptr(ptr)
    s.ptr = ptr
    return s
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (s *TShape) Free() {
    if s.instance != 0 {
        Shape_Free(s.instance)
        s.instance = 0
        s.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (s *TShape) Instance() uintptr {
    return s.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (s *TShape) UnsafeAddr() unsafe.Pointer {
    return s.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (s *TShape) IsValid() bool {
    return s.instance != 0
}

// TShapeClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TShapeClass() TClass {
    return Shape_StaticClassType()
}

// BringToFront
func (s *TShape) BringToFront() {
    Shape_BringToFront(s.instance)
}

// ClientToScreen
func (s *TShape) ClientToScreen(Point TPoint) TPoint {
    return Shape_ClientToScreen(s.instance, Point)
}

// ClientToParent
func (s *TShape) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return Shape_ClientToParent(s.instance, Point , CheckPtr(AParent))
}

// Dragging
func (s *TShape) Dragging() bool {
    return Shape_Dragging(s.instance)
}

// HasParent
func (s *TShape) HasParent() bool {
    return Shape_HasParent(s.instance)
}

// Hide
func (s *TShape) Hide() {
    Shape_Hide(s.instance)
}

// Invalidate
func (s *TShape) Invalidate() {
    Shape_Invalidate(s.instance)
}

// Perform
func (s *TShape) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return Shape_Perform(s.instance, Msg , WParam , LParam)
}

// Refresh
func (s *TShape) Refresh() {
    Shape_Refresh(s.instance)
}

// Repaint
func (s *TShape) Repaint() {
    Shape_Repaint(s.instance)
}

// ScreenToClient
func (s *TShape) ScreenToClient(Point TPoint) TPoint {
    return Shape_ScreenToClient(s.instance, Point)
}

// ParentToClient
func (s *TShape) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return Shape_ParentToClient(s.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (s *TShape) SendToBack() {
    Shape_SendToBack(s.instance)
}

// SetBounds
func (s *TShape) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    Shape_SetBounds(s.instance, ALeft , ATop , AWidth , AHeight)
}

// Show
func (s *TShape) Show() {
    Shape_Show(s.instance)
}

// Update
func (s *TShape) Update() {
    Shape_Update(s.instance)
}

// GetTextBuf
func (s *TShape) GetTextBuf(Buffer string, BufSize int32) int32 {
    return Shape_GetTextBuf(s.instance, Buffer , BufSize)
}

// GetTextLen
func (s *TShape) GetTextLen() int32 {
    return Shape_GetTextLen(s.instance)
}

// SetTextBuf
func (s *TShape) SetTextBuf(Buffer string) {
    Shape_SetTextBuf(s.instance, Buffer)
}

// FindComponent
func (s *TShape) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Shape_FindComponent(s.instance, AName))
}

// GetNamePath
func (s *TShape) GetNamePath() string {
    return Shape_GetNamePath(s.instance)
}

// Assign
func (s *TShape) Assign(Source IObject) {
    Shape_Assign(s.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (s *TShape) DisposeOf() {
    Shape_DisposeOf(s.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (s *TShape) ClassType() TClass {
    return Shape_ClassType(s.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (s *TShape) ClassName() string {
    return Shape_ClassName(s.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (s *TShape) InstanceSize() int32 {
    return Shape_InstanceSize(s.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (s *TShape) InheritsFrom(AClass TClass) bool {
    return Shape_InheritsFrom(s.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (s *TShape) Equals(Obj IObject) bool {
    return Shape_Equals(s.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (s *TShape) GetHashCode() int32 {
    return Shape_GetHashCode(s.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (s *TShape) ToString() string {
    return Shape_ToString(s.instance)
}

// Align
func (s *TShape) Align() TAlign {
    return Shape_GetAlign(s.instance)
}

// SetAlign
func (s *TShape) SetAlign(value TAlign) {
    Shape_SetAlign(s.instance, value)
}

// Anchors
func (s *TShape) Anchors() TAnchors {
    return Shape_GetAnchors(s.instance)
}

// SetAnchors
func (s *TShape) SetAnchors(value TAnchors) {
    Shape_SetAnchors(s.instance, value)
}

// Brush
func (s *TShape) Brush() *TBrush {
    return BrushFromInst(Shape_GetBrush(s.instance))
}

// SetBrush
func (s *TShape) SetBrush(value *TBrush) {
    Shape_SetBrush(s.instance, CheckPtr(value))
}

// DragCursor
func (s *TShape) DragCursor() TCursor {
    return Shape_GetDragCursor(s.instance)
}

// SetDragCursor
func (s *TShape) SetDragCursor(value TCursor) {
    Shape_SetDragCursor(s.instance, value)
}

// DragKind
func (s *TShape) DragKind() TDragKind {
    return Shape_GetDragKind(s.instance)
}

// SetDragKind
func (s *TShape) SetDragKind(value TDragKind) {
    Shape_SetDragKind(s.instance, value)
}

// DragMode
func (s *TShape) DragMode() TDragMode {
    return Shape_GetDragMode(s.instance)
}

// SetDragMode
func (s *TShape) SetDragMode(value TDragMode) {
    Shape_SetDragMode(s.instance, value)
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (s *TShape) Enabled() bool {
    return Shape_GetEnabled(s.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (s *TShape) SetEnabled(value bool) {
    Shape_SetEnabled(s.instance, value)
}

// ParentShowHint
func (s *TShape) ParentShowHint() bool {
    return Shape_GetParentShowHint(s.instance)
}

// SetParentShowHint
func (s *TShape) SetParentShowHint(value bool) {
    Shape_SetParentShowHint(s.instance, value)
}

// Pen
func (s *TShape) Pen() *TPen {
    return PenFromInst(Shape_GetPen(s.instance))
}

// SetPen
func (s *TShape) SetPen(value *TPen) {
    Shape_SetPen(s.instance, CheckPtr(value))
}

// Shape
func (s *TShape) Shape() TShapeType {
    return Shape_GetShape(s.instance)
}

// SetShape
func (s *TShape) SetShape(value TShapeType) {
    Shape_SetShape(s.instance, value)
}

// ShowHint
func (s *TShape) ShowHint() bool {
    return Shape_GetShowHint(s.instance)
}

// SetShowHint
func (s *TShape) SetShowHint(value bool) {
    Shape_SetShowHint(s.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (s *TShape) Visible() bool {
    return Shape_GetVisible(s.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (s *TShape) SetVisible(value bool) {
    Shape_SetVisible(s.instance, value)
}

// SetOnContextPopup
func (s *TShape) SetOnContextPopup(fn TContextPopupEvent) {
    Shape_SetOnContextPopup(s.instance, fn)
}

// SetOnDragDrop
func (s *TShape) SetOnDragDrop(fn TDragDropEvent) {
    Shape_SetOnDragDrop(s.instance, fn)
}

// SetOnDragOver
func (s *TShape) SetOnDragOver(fn TDragOverEvent) {
    Shape_SetOnDragOver(s.instance, fn)
}

// SetOnEndDock
func (s *TShape) SetOnEndDock(fn TEndDragEvent) {
    Shape_SetOnEndDock(s.instance, fn)
}

// SetOnEndDrag
func (s *TShape) SetOnEndDrag(fn TEndDragEvent) {
    Shape_SetOnEndDrag(s.instance, fn)
}

// SetOnMouseDown
func (s *TShape) SetOnMouseDown(fn TMouseEvent) {
    Shape_SetOnMouseDown(s.instance, fn)
}

// SetOnMouseEnter
func (s *TShape) SetOnMouseEnter(fn TNotifyEvent) {
    Shape_SetOnMouseEnter(s.instance, fn)
}

// SetOnMouseLeave
func (s *TShape) SetOnMouseLeave(fn TNotifyEvent) {
    Shape_SetOnMouseLeave(s.instance, fn)
}

// SetOnMouseMove
func (s *TShape) SetOnMouseMove(fn TMouseMoveEvent) {
    Shape_SetOnMouseMove(s.instance, fn)
}

// SetOnMouseUp
func (s *TShape) SetOnMouseUp(fn TMouseEvent) {
    Shape_SetOnMouseUp(s.instance, fn)
}

// SetOnStartDock
func (s *TShape) SetOnStartDock(fn TStartDockEvent) {
    Shape_SetOnStartDock(s.instance, fn)
}

// Action
func (s *TShape) Action() *TAction {
    return ActionFromInst(Shape_GetAction(s.instance))
}

// SetAction
func (s *TShape) SetAction(value IComponent) {
    Shape_SetAction(s.instance, CheckPtr(value))
}

// BiDiMode
func (s *TShape) BiDiMode() TBiDiMode {
    return Shape_GetBiDiMode(s.instance)
}

// SetBiDiMode
func (s *TShape) SetBiDiMode(value TBiDiMode) {
    Shape_SetBiDiMode(s.instance, value)
}

// BoundsRect
func (s *TShape) BoundsRect() TRect {
    return Shape_GetBoundsRect(s.instance)
}

// SetBoundsRect
func (s *TShape) SetBoundsRect(value TRect) {
    Shape_SetBoundsRect(s.instance, value)
}

// ClientHeight
func (s *TShape) ClientHeight() int32 {
    return Shape_GetClientHeight(s.instance)
}

// SetClientHeight
func (s *TShape) SetClientHeight(value int32) {
    Shape_SetClientHeight(s.instance, value)
}

// ClientOrigin
func (s *TShape) ClientOrigin() TPoint {
    return Shape_GetClientOrigin(s.instance)
}

// ClientRect
func (s *TShape) ClientRect() TRect {
    return Shape_GetClientRect(s.instance)
}

// ClientWidth
func (s *TShape) ClientWidth() int32 {
    return Shape_GetClientWidth(s.instance)
}

// SetClientWidth
func (s *TShape) SetClientWidth(value int32) {
    Shape_SetClientWidth(s.instance, value)
}

// ControlState
func (s *TShape) ControlState() TControlState {
    return Shape_GetControlState(s.instance)
}

// SetControlState
func (s *TShape) SetControlState(value TControlState) {
    Shape_SetControlState(s.instance, value)
}

// ControlStyle
func (s *TShape) ControlStyle() TControlStyle {
    return Shape_GetControlStyle(s.instance)
}

// SetControlStyle
func (s *TShape) SetControlStyle(value TControlStyle) {
    Shape_SetControlStyle(s.instance, value)
}

// ExplicitLeft
func (s *TShape) ExplicitLeft() int32 {
    return Shape_GetExplicitLeft(s.instance)
}

// ExplicitTop
func (s *TShape) ExplicitTop() int32 {
    return Shape_GetExplicitTop(s.instance)
}

// ExplicitWidth
func (s *TShape) ExplicitWidth() int32 {
    return Shape_GetExplicitWidth(s.instance)
}

// ExplicitHeight
func (s *TShape) ExplicitHeight() int32 {
    return Shape_GetExplicitHeight(s.instance)
}

// Floating
func (s *TShape) Floating() bool {
    return Shape_GetFloating(s.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (s *TShape) Parent() *TWinControl {
    return WinControlFromInst(Shape_GetParent(s.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (s *TShape) SetParent(value IWinControl) {
    Shape_SetParent(s.instance, CheckPtr(value))
}

// StyleElements
func (s *TShape) StyleElements() TStyleElements {
    return Shape_GetStyleElements(s.instance)
}

// SetStyleElements
func (s *TShape) SetStyleElements(value TStyleElements) {
    Shape_SetStyleElements(s.instance, value)
}

// AlignWithMargins
func (s *TShape) AlignWithMargins() bool {
    return Shape_GetAlignWithMargins(s.instance)
}

// SetAlignWithMargins
func (s *TShape) SetAlignWithMargins(value bool) {
    Shape_SetAlignWithMargins(s.instance, value)
}

// Left
func (s *TShape) Left() int32 {
    return Shape_GetLeft(s.instance)
}

// SetLeft
func (s *TShape) SetLeft(value int32) {
    Shape_SetLeft(s.instance, value)
}

// Top
func (s *TShape) Top() int32 {
    return Shape_GetTop(s.instance)
}

// SetTop
func (s *TShape) SetTop(value int32) {
    Shape_SetTop(s.instance, value)
}

// Width
func (s *TShape) Width() int32 {
    return Shape_GetWidth(s.instance)
}

// SetWidth
func (s *TShape) SetWidth(value int32) {
    Shape_SetWidth(s.instance, value)
}

// Height
func (s *TShape) Height() int32 {
    return Shape_GetHeight(s.instance)
}

// SetHeight
func (s *TShape) SetHeight(value int32) {
    Shape_SetHeight(s.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (s *TShape) Cursor() TCursor {
    return Shape_GetCursor(s.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (s *TShape) SetCursor(value TCursor) {
    Shape_SetCursor(s.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (s *TShape) Hint() string {
    return Shape_GetHint(s.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (s *TShape) SetHint(value string) {
    Shape_SetHint(s.instance, value)
}

// Margins
func (s *TShape) Margins() *TMargins {
    return MarginsFromInst(Shape_GetMargins(s.instance))
}

// SetMargins
func (s *TShape) SetMargins(value *TMargins) {
    Shape_SetMargins(s.instance, CheckPtr(value))
}

// CustomHint
func (s *TShape) CustomHint() *TCustomHint {
    return CustomHintFromInst(Shape_GetCustomHint(s.instance))
}

// SetCustomHint
func (s *TShape) SetCustomHint(value IComponent) {
    Shape_SetCustomHint(s.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (s *TShape) ComponentCount() int32 {
    return Shape_GetComponentCount(s.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (s *TShape) ComponentIndex() int32 {
    return Shape_GetComponentIndex(s.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (s *TShape) SetComponentIndex(value int32) {
    Shape_SetComponentIndex(s.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (s *TShape) Owner() *TComponent {
    return ComponentFromInst(Shape_GetOwner(s.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (s *TShape) Name() string {
    return Shape_GetName(s.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (s *TShape) SetName(value string) {
    Shape_SetName(s.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (s *TShape) Tag() int {
    return Shape_GetTag(s.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (s *TShape) SetTag(value int) {
    Shape_SetTag(s.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (s *TShape) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Shape_GetComponents(s.instance, AIndex))
}

