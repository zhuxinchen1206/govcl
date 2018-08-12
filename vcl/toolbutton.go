
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

type TToolButton struct {
    IControl
    instance uintptr
    // 特殊情况下使用，主要应对Go的GC问题，与VCL没有太多关系。
    ptr unsafe.Pointer
}

// NewToolButton
// CN: 创建一个新的对象。
// EN: Create a new object.
func NewToolButton(owner IComponent) *TToolButton {
    t := new(TToolButton)
    t.instance = ToolButton_Create(CheckPtr(owner))
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolButtonFromInst
// CN: 新建一个对象来自已经存在的对象实例指针。
// EN: Create a new object from an existing object instance pointer.
func ToolButtonFromInst(inst uintptr) *TToolButton {
    t := new(TToolButton)
    t.instance = inst
    t.ptr = unsafe.Pointer(inst)
    return t
}

// ToolButtonFromObj
// CN: 新建一个对象来自已经存在的对象实例。
// EN: Create a new object from an existing object instance.
func ToolButtonFromObj(obj IObject) *TToolButton {
    t := new(TToolButton)
    t.instance = CheckPtr(obj)
    t.ptr = unsafe.Pointer(t.instance)
    return t
}

// ToolButtonFromUnsafePointer
// CN: 新建一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
// EN: Create a new object from an unsecure address. Note: Using this function may cause some unclear situations and be used with caution..
func ToolButtonFromUnsafePointer(ptr unsafe.Pointer) *TToolButton {
    t := new(TToolButton)
    t.instance = uintptr(ptr)
    t.ptr = ptr
    return t
}

// Free 
// CN: 释放对象。
// EN: Free object.
func (t *TToolButton) Free() {
    if t.instance != 0 {
        ToolButton_Free(t.instance)
        t.instance = 0
        t.ptr = unsafe.Pointer(uintptr(0))
    }
}

// Instance 
// CN: 返回对象实例指针。
// EN: Return object instance pointer.
func (t *TToolButton) Instance() uintptr {
    return t.instance
}

// UnsafeAddr 
// CN: 获取一个不安全的地址。
// EN: Get an unsafe address.
func (t *TToolButton) UnsafeAddr() unsafe.Pointer {
    return t.ptr
}

// IsValid 
// CN: 检测地址是否为空。
// EN: Check if the address is empty.
func (t *TToolButton) IsValid() bool {
    return t.instance != 0
}

// TToolButtonClass
// CN: 获取类信息指针。
// EN: Get class information pointer.
func TToolButtonClass() TClass {
    return ToolButton_StaticClassType()
}

// CheckMenuDropdown
func (t *TToolButton) CheckMenuDropdown() bool {
    return ToolButton_CheckMenuDropdown(t.instance)
}

// Click
func (t *TToolButton) Click() {
    ToolButton_Click(t.instance)
}

// SetBounds
func (t *TToolButton) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    ToolButton_SetBounds(t.instance, ALeft , ATop , AWidth , AHeight)
}

// BringToFront
func (t *TToolButton) BringToFront() {
    ToolButton_BringToFront(t.instance)
}

// ClientToScreen
func (t *TToolButton) ClientToScreen(Point TPoint) TPoint {
    return ToolButton_ClientToScreen(t.instance, Point)
}

// ClientToParent
func (t *TToolButton) ClientToParent(Point TPoint, AParent IWinControl) TPoint {
    return ToolButton_ClientToParent(t.instance, Point , CheckPtr(AParent))
}

// Dragging
func (t *TToolButton) Dragging() bool {
    return ToolButton_Dragging(t.instance)
}

// HasParent
func (t *TToolButton) HasParent() bool {
    return ToolButton_HasParent(t.instance)
}

// Hide
func (t *TToolButton) Hide() {
    ToolButton_Hide(t.instance)
}

// Invalidate
func (t *TToolButton) Invalidate() {
    ToolButton_Invalidate(t.instance)
}

// Perform
func (t *TToolButton) Perform(Msg uint32, WParam uintptr, LParam int) int {
    return ToolButton_Perform(t.instance, Msg , WParam , LParam)
}

// Refresh
func (t *TToolButton) Refresh() {
    ToolButton_Refresh(t.instance)
}

// Repaint
func (t *TToolButton) Repaint() {
    ToolButton_Repaint(t.instance)
}

// ScreenToClient
func (t *TToolButton) ScreenToClient(Point TPoint) TPoint {
    return ToolButton_ScreenToClient(t.instance, Point)
}

// ParentToClient
func (t *TToolButton) ParentToClient(Point TPoint, AParent IWinControl) TPoint {
    return ToolButton_ParentToClient(t.instance, Point , CheckPtr(AParent))
}

// SendToBack
func (t *TToolButton) SendToBack() {
    ToolButton_SendToBack(t.instance)
}

// Show
func (t *TToolButton) Show() {
    ToolButton_Show(t.instance)
}

// Update
func (t *TToolButton) Update() {
    ToolButton_Update(t.instance)
}

// GetTextBuf
func (t *TToolButton) GetTextBuf(Buffer string, BufSize int32) int32 {
    return ToolButton_GetTextBuf(t.instance, Buffer , BufSize)
}

// GetTextLen
func (t *TToolButton) GetTextLen() int32 {
    return ToolButton_GetTextLen(t.instance)
}

// SetTextBuf
func (t *TToolButton) SetTextBuf(Buffer string) {
    ToolButton_SetTextBuf(t.instance, Buffer)
}

// FindComponent
func (t *TToolButton) FindComponent(AName string) *TComponent {
    return ComponentFromInst(ToolButton_FindComponent(t.instance, AName))
}

// GetNamePath
func (t *TToolButton) GetNamePath() string {
    return ToolButton_GetNamePath(t.instance)
}

// Assign
func (t *TToolButton) Assign(Source IObject) {
    ToolButton_Assign(t.instance, CheckPtr(Source))
}

// DisposeOf
// CN: 丢弃当前对象。
// EN: Discard the current object.
func (t *TToolButton) DisposeOf() {
    ToolButton_DisposeOf(t.instance)
}

// ClassType
// CN: 获取类的类型信息。
// EN: Get class type information.
func (t *TToolButton) ClassType() TClass {
    return ToolButton_ClassType(t.instance)
}

// ClassName
// CN: 获取当前对象类名称。
// EN: Get the current object class name.
func (t *TToolButton) ClassName() string {
    return ToolButton_ClassName(t.instance)
}

// InstanceSize
// CN: 获取当前对象实例大小。
// EN: Get the current object instance size.
func (t *TToolButton) InstanceSize() int32 {
    return ToolButton_InstanceSize(t.instance)
}

// InheritsFrom
// CN: 判断当前类是否继承自指定类。
// EN: Determine whether the current class inherits from the specified class.
func (t *TToolButton) InheritsFrom(AClass TClass) bool {
    return ToolButton_InheritsFrom(t.instance, AClass)
}

// Equals
// CN: 与一个对象进行比较。
// EN: Compare with an object.
func (t *TToolButton) Equals(Obj IObject) bool {
    return ToolButton_Equals(t.instance, CheckPtr(Obj))
}

// GetHashCode
// CN: 获取类的哈希值。
// EN: Get the hash value of the class.
func (t *TToolButton) GetHashCode() int32 {
    return ToolButton_GetHashCode(t.instance)
}

// ToString
// CN: 文本类信息。
// EN: Text information.
func (t *TToolButton) ToString() string {
    return ToolButton_ToString(t.instance)
}

// Index
func (t *TToolButton) Index() int32 {
    return ToolButton_GetIndex(t.instance)
}

// Action
func (t *TToolButton) Action() *TAction {
    return ActionFromInst(ToolButton_GetAction(t.instance))
}

// SetAction
func (t *TToolButton) SetAction(value IComponent) {
    ToolButton_SetAction(t.instance, CheckPtr(value))
}

// AllowAllUp
func (t *TToolButton) AllowAllUp() bool {
    return ToolButton_GetAllowAllUp(t.instance)
}

// SetAllowAllUp
func (t *TToolButton) SetAllowAllUp(value bool) {
    ToolButton_SetAllowAllUp(t.instance, value)
}

// AutoSize
func (t *TToolButton) AutoSize() bool {
    return ToolButton_GetAutoSize(t.instance)
}

// SetAutoSize
func (t *TToolButton) SetAutoSize(value bool) {
    ToolButton_SetAutoSize(t.instance, value)
}

// Caption
// CN: 获取控件标题。
// EN: Get the control title.
func (t *TToolButton) Caption() string {
    return ToolButton_GetCaption(t.instance)
}

// SetCaption
// CN: 设置控件标题。
// EN: Set the control title.
func (t *TToolButton) SetCaption(value string) {
    ToolButton_SetCaption(t.instance, value)
}

// Down
func (t *TToolButton) Down() bool {
    return ToolButton_GetDown(t.instance)
}

// SetDown
func (t *TToolButton) SetDown(value bool) {
    ToolButton_SetDown(t.instance, value)
}

// DragCursor
func (t *TToolButton) DragCursor() TCursor {
    return ToolButton_GetDragCursor(t.instance)
}

// SetDragCursor
func (t *TToolButton) SetDragCursor(value TCursor) {
    ToolButton_SetDragCursor(t.instance, value)
}

// DragKind
func (t *TToolButton) DragKind() TDragKind {
    return ToolButton_GetDragKind(t.instance)
}

// SetDragKind
func (t *TToolButton) SetDragKind(value TDragKind) {
    ToolButton_SetDragKind(t.instance, value)
}

// DragMode
func (t *TToolButton) DragMode() TDragMode {
    return ToolButton_GetDragMode(t.instance)
}

// SetDragMode
func (t *TToolButton) SetDragMode(value TDragMode) {
    ToolButton_SetDragMode(t.instance, value)
}

// DropdownMenu
func (t *TToolButton) DropdownMenu() *TPopupMenu {
    return PopupMenuFromInst(ToolButton_GetDropdownMenu(t.instance))
}

// SetDropdownMenu
func (t *TToolButton) SetDropdownMenu(value IComponent) {
    ToolButton_SetDropdownMenu(t.instance, CheckPtr(value))
}

// Enabled
// CN: 获取控件启用。
// EN: Get the control enabled.
func (t *TToolButton) Enabled() bool {
    return ToolButton_GetEnabled(t.instance)
}

// SetEnabled
// CN: 设置控件启用。
// EN: Set the control enabled.
func (t *TToolButton) SetEnabled(value bool) {
    ToolButton_SetEnabled(t.instance, value)
}

// EnableDropdown
func (t *TToolButton) EnableDropdown() bool {
    return ToolButton_GetEnableDropdown(t.instance)
}

// SetEnableDropdown
func (t *TToolButton) SetEnableDropdown(value bool) {
    ToolButton_SetEnableDropdown(t.instance, value)
}

// Grouped
func (t *TToolButton) Grouped() bool {
    return ToolButton_GetGrouped(t.instance)
}

// SetGrouped
func (t *TToolButton) SetGrouped(value bool) {
    ToolButton_SetGrouped(t.instance, value)
}

// Height
func (t *TToolButton) Height() int32 {
    return ToolButton_GetHeight(t.instance)
}

// SetHeight
func (t *TToolButton) SetHeight(value int32) {
    ToolButton_SetHeight(t.instance, value)
}

// ImageIndex
func (t *TToolButton) ImageIndex() int32 {
    return ToolButton_GetImageIndex(t.instance)
}

// SetImageIndex
func (t *TToolButton) SetImageIndex(value int32) {
    ToolButton_SetImageIndex(t.instance, value)
}

// Indeterminate
func (t *TToolButton) Indeterminate() bool {
    return ToolButton_GetIndeterminate(t.instance)
}

// SetIndeterminate
func (t *TToolButton) SetIndeterminate(value bool) {
    ToolButton_SetIndeterminate(t.instance, value)
}

// Marked
func (t *TToolButton) Marked() bool {
    return ToolButton_GetMarked(t.instance)
}

// SetMarked
func (t *TToolButton) SetMarked(value bool) {
    ToolButton_SetMarked(t.instance, value)
}

// ParentShowHint
func (t *TToolButton) ParentShowHint() bool {
    return ToolButton_GetParentShowHint(t.instance)
}

// SetParentShowHint
func (t *TToolButton) SetParentShowHint(value bool) {
    ToolButton_SetParentShowHint(t.instance, value)
}

// PopupMenu
func (t *TToolButton) PopupMenu() *TPopupMenu {
    return PopupMenuFromInst(ToolButton_GetPopupMenu(t.instance))
}

// SetPopupMenu
func (t *TToolButton) SetPopupMenu(value IComponent) {
    ToolButton_SetPopupMenu(t.instance, CheckPtr(value))
}

// Wrap
func (t *TToolButton) Wrap() bool {
    return ToolButton_GetWrap(t.instance)
}

// SetWrap
func (t *TToolButton) SetWrap(value bool) {
    ToolButton_SetWrap(t.instance, value)
}

// ShowHint
func (t *TToolButton) ShowHint() bool {
    return ToolButton_GetShowHint(t.instance)
}

// SetShowHint
func (t *TToolButton) SetShowHint(value bool) {
    ToolButton_SetShowHint(t.instance, value)
}

// Style
func (t *TToolButton) Style() TToolButtonStyle {
    return ToolButton_GetStyle(t.instance)
}

// SetStyle
func (t *TToolButton) SetStyle(value TToolButtonStyle) {
    ToolButton_SetStyle(t.instance, value)
}

// Visible
// CN: 获取控件可视。
// EN: Get the control visible.
func (t *TToolButton) Visible() bool {
    return ToolButton_GetVisible(t.instance)
}

// SetVisible
// CN: 设置控件可视。
// EN: Set the control visible.
func (t *TToolButton) SetVisible(value bool) {
    ToolButton_SetVisible(t.instance, value)
}

// Width
func (t *TToolButton) Width() int32 {
    return ToolButton_GetWidth(t.instance)
}

// SetWidth
func (t *TToolButton) SetWidth(value int32) {
    ToolButton_SetWidth(t.instance, value)
}

// SetOnClick
// CN: 设置控件单击事件。
// EN: Set control click event.
func (t *TToolButton) SetOnClick(fn TNotifyEvent) {
    ToolButton_SetOnClick(t.instance, fn)
}

// SetOnContextPopup
func (t *TToolButton) SetOnContextPopup(fn TContextPopupEvent) {
    ToolButton_SetOnContextPopup(t.instance, fn)
}

// SetOnDragDrop
func (t *TToolButton) SetOnDragDrop(fn TDragDropEvent) {
    ToolButton_SetOnDragDrop(t.instance, fn)
}

// SetOnDragOver
func (t *TToolButton) SetOnDragOver(fn TDragOverEvent) {
    ToolButton_SetOnDragOver(t.instance, fn)
}

// SetOnEndDock
func (t *TToolButton) SetOnEndDock(fn TEndDragEvent) {
    ToolButton_SetOnEndDock(t.instance, fn)
}

// SetOnEndDrag
func (t *TToolButton) SetOnEndDrag(fn TEndDragEvent) {
    ToolButton_SetOnEndDrag(t.instance, fn)
}

// SetOnMouseDown
func (t *TToolButton) SetOnMouseDown(fn TMouseEvent) {
    ToolButton_SetOnMouseDown(t.instance, fn)
}

// SetOnMouseEnter
func (t *TToolButton) SetOnMouseEnter(fn TNotifyEvent) {
    ToolButton_SetOnMouseEnter(t.instance, fn)
}

// SetOnMouseLeave
func (t *TToolButton) SetOnMouseLeave(fn TNotifyEvent) {
    ToolButton_SetOnMouseLeave(t.instance, fn)
}

// SetOnMouseMove
func (t *TToolButton) SetOnMouseMove(fn TMouseMoveEvent) {
    ToolButton_SetOnMouseMove(t.instance, fn)
}

// SetOnMouseUp
func (t *TToolButton) SetOnMouseUp(fn TMouseEvent) {
    ToolButton_SetOnMouseUp(t.instance, fn)
}

// SetOnStartDock
func (t *TToolButton) SetOnStartDock(fn TStartDockEvent) {
    ToolButton_SetOnStartDock(t.instance, fn)
}

// Align
func (t *TToolButton) Align() TAlign {
    return ToolButton_GetAlign(t.instance)
}

// SetAlign
func (t *TToolButton) SetAlign(value TAlign) {
    ToolButton_SetAlign(t.instance, value)
}

// Anchors
func (t *TToolButton) Anchors() TAnchors {
    return ToolButton_GetAnchors(t.instance)
}

// SetAnchors
func (t *TToolButton) SetAnchors(value TAnchors) {
    ToolButton_SetAnchors(t.instance, value)
}

// BiDiMode
func (t *TToolButton) BiDiMode() TBiDiMode {
    return ToolButton_GetBiDiMode(t.instance)
}

// SetBiDiMode
func (t *TToolButton) SetBiDiMode(value TBiDiMode) {
    ToolButton_SetBiDiMode(t.instance, value)
}

// BoundsRect
func (t *TToolButton) BoundsRect() TRect {
    return ToolButton_GetBoundsRect(t.instance)
}

// SetBoundsRect
func (t *TToolButton) SetBoundsRect(value TRect) {
    ToolButton_SetBoundsRect(t.instance, value)
}

// ClientHeight
func (t *TToolButton) ClientHeight() int32 {
    return ToolButton_GetClientHeight(t.instance)
}

// SetClientHeight
func (t *TToolButton) SetClientHeight(value int32) {
    ToolButton_SetClientHeight(t.instance, value)
}

// ClientOrigin
func (t *TToolButton) ClientOrigin() TPoint {
    return ToolButton_GetClientOrigin(t.instance)
}

// ClientRect
func (t *TToolButton) ClientRect() TRect {
    return ToolButton_GetClientRect(t.instance)
}

// ClientWidth
func (t *TToolButton) ClientWidth() int32 {
    return ToolButton_GetClientWidth(t.instance)
}

// SetClientWidth
func (t *TToolButton) SetClientWidth(value int32) {
    ToolButton_SetClientWidth(t.instance, value)
}

// ControlState
func (t *TToolButton) ControlState() TControlState {
    return ToolButton_GetControlState(t.instance)
}

// SetControlState
func (t *TToolButton) SetControlState(value TControlState) {
    ToolButton_SetControlState(t.instance, value)
}

// ControlStyle
func (t *TToolButton) ControlStyle() TControlStyle {
    return ToolButton_GetControlStyle(t.instance)
}

// SetControlStyle
func (t *TToolButton) SetControlStyle(value TControlStyle) {
    ToolButton_SetControlStyle(t.instance, value)
}

// ExplicitLeft
func (t *TToolButton) ExplicitLeft() int32 {
    return ToolButton_GetExplicitLeft(t.instance)
}

// ExplicitTop
func (t *TToolButton) ExplicitTop() int32 {
    return ToolButton_GetExplicitTop(t.instance)
}

// ExplicitWidth
func (t *TToolButton) ExplicitWidth() int32 {
    return ToolButton_GetExplicitWidth(t.instance)
}

// ExplicitHeight
func (t *TToolButton) ExplicitHeight() int32 {
    return ToolButton_GetExplicitHeight(t.instance)
}

// Floating
func (t *TToolButton) Floating() bool {
    return ToolButton_GetFloating(t.instance)
}

// Parent
// CN: 获取控件父容器。
// EN: Get control parent container.
func (t *TToolButton) Parent() *TWinControl {
    return WinControlFromInst(ToolButton_GetParent(t.instance))
}

// SetParent
// CN: 设置控件父容器。
// EN: Set control parent container.
func (t *TToolButton) SetParent(value IWinControl) {
    ToolButton_SetParent(t.instance, CheckPtr(value))
}

// StyleElements
func (t *TToolButton) StyleElements() TStyleElements {
    return ToolButton_GetStyleElements(t.instance)
}

// SetStyleElements
func (t *TToolButton) SetStyleElements(value TStyleElements) {
    ToolButton_SetStyleElements(t.instance, value)
}

// AlignWithMargins
func (t *TToolButton) AlignWithMargins() bool {
    return ToolButton_GetAlignWithMargins(t.instance)
}

// SetAlignWithMargins
func (t *TToolButton) SetAlignWithMargins(value bool) {
    ToolButton_SetAlignWithMargins(t.instance, value)
}

// Left
func (t *TToolButton) Left() int32 {
    return ToolButton_GetLeft(t.instance)
}

// SetLeft
func (t *TToolButton) SetLeft(value int32) {
    ToolButton_SetLeft(t.instance, value)
}

// Top
func (t *TToolButton) Top() int32 {
    return ToolButton_GetTop(t.instance)
}

// SetTop
func (t *TToolButton) SetTop(value int32) {
    ToolButton_SetTop(t.instance, value)
}

// Cursor
// CN: 获取控件光标。
// EN: Get control cursor.
func (t *TToolButton) Cursor() TCursor {
    return ToolButton_GetCursor(t.instance)
}

// SetCursor
// CN: 设置控件光标。
// EN: Set control cursor.
func (t *TToolButton) SetCursor(value TCursor) {
    ToolButton_SetCursor(t.instance, value)
}

// Hint
// CN: 获取组件鼠标提示。
// EN: Get component mouse hints.
func (t *TToolButton) Hint() string {
    return ToolButton_GetHint(t.instance)
}

// SetHint
// CN: 设置组件鼠标提示。
// EN: Set component mouse hints.
func (t *TToolButton) SetHint(value string) {
    ToolButton_SetHint(t.instance, value)
}

// Margins
func (t *TToolButton) Margins() *TMargins {
    return MarginsFromInst(ToolButton_GetMargins(t.instance))
}

// SetMargins
func (t *TToolButton) SetMargins(value *TMargins) {
    ToolButton_SetMargins(t.instance, CheckPtr(value))
}

// CustomHint
func (t *TToolButton) CustomHint() *TCustomHint {
    return CustomHintFromInst(ToolButton_GetCustomHint(t.instance))
}

// SetCustomHint
func (t *TToolButton) SetCustomHint(value IComponent) {
    ToolButton_SetCustomHint(t.instance, CheckPtr(value))
}

// ComponentCount
// CN: 获取组件总数。
// EN: Get the total number of components.
func (t *TToolButton) ComponentCount() int32 {
    return ToolButton_GetComponentCount(t.instance)
}

// ComponentIndex
// CN: 获取组件索引。
// EN: Get component index.
func (t *TToolButton) ComponentIndex() int32 {
    return ToolButton_GetComponentIndex(t.instance)
}

// SetComponentIndex
// CN: 设置组件索引。
// EN: Set component index.
func (t *TToolButton) SetComponentIndex(value int32) {
    ToolButton_SetComponentIndex(t.instance, value)
}

// Owner
// CN: 获取组件所有者。
// EN: Get component owner.
func (t *TToolButton) Owner() *TComponent {
    return ComponentFromInst(ToolButton_GetOwner(t.instance))
}

// Name
// CN: 获取组件名称。
// EN: Get the component name.
func (t *TToolButton) Name() string {
    return ToolButton_GetName(t.instance)
}

// SetName
// CN: 设置组件名称。
// EN: Set the component name.
func (t *TToolButton) SetName(value string) {
    ToolButton_SetName(t.instance, value)
}

// Tag
// CN: 获取对象标记。
// EN: Get the control tag.
func (t *TToolButton) Tag() int {
    return ToolButton_GetTag(t.instance)
}

// SetTag
// CN: 设置对象标记。
// EN: Set the control tag.
func (t *TToolButton) SetTag(value int) {
    ToolButton_SetTag(t.instance, value)
}

// Components
// CN: 获取指定索引组件。
// EN: Get the specified index component.
func (t *TToolButton) Components(AIndex int32) *TComponent {
    return ComponentFromInst(ToolButton_GetComponents(t.instance, AIndex))
}

