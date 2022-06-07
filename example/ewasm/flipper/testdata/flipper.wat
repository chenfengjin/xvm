(module
  (type (;0;) (func (param i32 i32 i32) (result i32)))
  (type (;1;) (func (param i32 i32 i32)))
  (type (;2;) (func (param i32 i32)))
  (type (;3;) (func (param i32) (result i32)))
  (type (;4;) (func))
  (type (;5;) (func (result i32)))
  (import "env" "memory" (memory (;0;) 16 16))
  (import "seal0" "seal_get_storage" (func $seal_get_storage (type 0)))
  (import "seal0" "seal_set_storage" (func $seal_set_storage (type 1)))
  (import "seal0" "seal_input" (func $seal_input (type 2)))
  (import "seal0" "seal_return" (func $seal_return (type 1)))
  (import "seal0" "seal_value_transferred" (func $seal_value_transferred (type 2)))
  (func $__malloc (type 3) (param i32) (result i32)
    (local i32 i32 i32 i32)
    i32.const 65536
    local.set 1
    loop (result i32)  ;; label = @1
      block  ;; label = @2
        local.get 1
        i32.load offset=12
        br_if 0 (;@2;)
        local.get 1
        i32.load offset=8
        local.tee 2
        local.get 0
        i32.lt_u
        br_if 0 (;@2;)
        block  ;; label = @3
          local.get 2
          local.get 0
          i32.const 7
          i32.add
          i32.const -8
          i32.and
          local.tee 3
          i32.sub
          local.tee 2
          i32.const 24
          i32.lt_u
          br_if 0 (;@3;)
          local.get 1
          local.get 3
          i32.add
          i32.const 16
          i32.add
          local.tee 0
          local.get 1
          i32.load
          local.tee 4
          i32.store
          block  ;; label = @4
            local.get 4
            i32.eqz
            br_if 0 (;@4;)
            local.get 4
            local.get 0
            i32.store offset=4
          end
          local.get 0
          local.get 2
          i32.const -16
          i32.add
          i32.store offset=8
          local.get 0
          i32.const 0
          i32.store offset=12
          local.get 0
          local.get 1
          i32.store offset=4
          local.get 1
          local.get 0
          i32.store
          local.get 1
          local.get 3
          i32.store offset=8
        end
        local.get 1
        i32.const 1
        i32.store offset=12
        local.get 1
        i32.const 16
        i32.add
        return
      end
      local.get 1
      i32.load
      local.set 1
      br 0 (;@1;)
    end)
  (func $__init_heap (type 4)
    i32.const 0
    i32.const 0
    i32.store offset=65536
    i32.const 0
    i32.const 0
    i32.store offset=65540
    i32.const 0
    i32.const 0
    i32.store offset=65548
    i32.const 0
    memory.size
    i32.const 16
    i32.shl
    i32.const -65552
    i32.add
    i32.store offset=65544)
  (func $flipper::flipper::function::flip (type 5) (result i32)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 0
    global.set $__stack_pointer
    i32.const 0
    i32.const 1
    i32.store offset=8
    local.get 0
    local.tee 1
    i32.const 8
    i32.add
    i32.const 24
    i32.add
    i64.const 0
    i64.store
    local.get 1
    i64.const 0
    i64.store offset=24
    local.get 1
    i64.const 0
    i64.store offset=16
    local.get 1
    i64.const 0
    i64.store offset=8
    i32.const 0
    local.set 2
    block  ;; label = @1
      local.get 1
      i32.const 8
      i32.add
      i32.const 16
      i32.const 8
      call $seal_get_storage
      br_if 0 (;@1;)
      i32.const 0
      i32.load8_u offset=16
      local.set 2
    end
    local.get 0
    i32.const -32
    i32.add
    local.tee 0
    global.set $__stack_pointer
    local.get 0
    i32.const 24
    i32.add
    i64.const 0
    i64.store
    local.get 0
    i64.const 0
    i64.store offset=16
    local.get 0
    i64.const 0
    i64.store offset=8
    local.get 0
    i64.const 0
    i64.store
    local.get 1
    local.get 2
    i32.const -1
    i32.xor
    i32.const 1
    i32.and
    i32.store8 offset=47
    local.get 0
    local.get 1
    i32.const 47
    i32.add
    i32.const 1
    call $seal_set_storage
    local.get 1
    i32.const 48
    i32.add
    global.set $__stack_pointer
    i32.const 0)
  (func $deploy (type 4)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 48
    i32.sub
    local.tee 0
    global.set $__stack_pointer
    call $__init_heap
    i32.const 0
    i32.const 32768
    i32.store offset=8
    i32.const 16
    i32.const 8
    call $seal_input
    i32.const 0
    i32.const 0
    i32.load offset=8
    local.tee 1
    i32.store offset=4
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            local.get 1
            i32.const 3
            i32.le_u
            br_if 0 (;@4;)
            i32.const 0
            i32.const 0
            i32.load offset=16
            local.tee 2
            i32.store
            local.get 2
            i32.const 444473080
            i32.ne
            br_if 0 (;@4;)
            i32.const 21
            local.get 1
            i32.const -4
            i32.add
            i32.const 20
            i32.add
            local.tee 1
            i32.gt_u
            br_if 1 (;@3;)
            i32.const 21
            local.get 1
            i32.ne
            br_if 2 (;@2;)
            i32.const 0
            i32.load8_u offset=20
            local.set 1
            local.get 0
            i32.const 32
            i32.add
            i64.const 0
            i64.store
            local.get 0
            i64.const 0
            i64.store offset=24
            local.get 0
            i64.const 0
            i64.store offset=16
            local.get 0
            i64.const 0
            i64.store offset=8
            local.get 0
            local.get 1
            i32.const 255
            i32.and
            i32.const 1
            i32.eq
            i32.store8 offset=47
            local.get 0
            i32.const 8
            i32.add
            local.get 0
            i32.const 47
            i32.add
            i32.const 1
            call $seal_set_storage
            i32.const 1
            br_if 3 (;@1;)
            unreachable
            unreachable
          end
          unreachable
          unreachable
        end
        unreachable
        unreachable
      end
      unreachable
      unreachable
    end
    i32.const 0
    i32.const 0
    i32.const 0
    call $seal_return
    unreachable)
  (func $call (type 4)
    (local i32 i32 i32)
    global.get $__stack_pointer
    i32.const 64
    i32.sub
    local.tee 0
    global.set $__stack_pointer
    local.get 0
    i32.const 16
    i32.store offset=4
    local.get 0
    i32.const 8
    i32.add
    local.get 0
    i32.const 4
    i32.add
    call $seal_value_transferred
    block  ;; label = @1
      block  ;; label = @2
        block  ;; label = @3
          block  ;; label = @4
            block  ;; label = @5
              block  ;; label = @6
                local.get 0
                i64.load offset=8
                local.get 0
                i32.const 16
                i32.add
                i64.load
                i64.or
                i64.const 0
                i64.ne
                br_if 0 (;@6;)
                call $__init_heap
                i32.const 0
                i32.const 32768
                i32.store offset=8
                i32.const 16
                i32.const 8
                call $seal_input
                i32.const 0
                i32.const 0
                i32.load offset=8
                local.tee 1
                i32.store offset=4
                local.get 1
                i32.const 3
                i32.le_u
                br_if 1 (;@5;)
                i32.const 0
                i32.const 0
                i32.load offset=16
                local.tee 2
                i32.store
                local.get 1
                i32.const -4
                i32.add
                local.set 1
                block  ;; label = @7
                  local.get 2
                  i32.const 1021725805
                  i32.eq
                  br_if 0 (;@7;)
                  local.get 2
                  i32.const -1443896115
                  i32.ne
                  br_if 2 (;@5;)
                  local.get 1
                  br_if 3 (;@4;)
                  call $flipper::flipper::function::flip
                  i32.eqz
                  br_if 4 (;@3;)
                  unreachable
                  unreachable
                end
                local.get 1
                br_if 4 (;@2;)
                i32.const 0
                local.set 1
                i32.const 0
                i32.const 1
                i32.store offset=8
                local.get 0
                i32.const 56
                i32.add
                i64.const 0
                i64.store
                local.get 0
                i64.const 0
                i64.store offset=48
                local.get 0
                i64.const 0
                i64.store offset=40
                local.get 0
                i64.const 0
                i64.store offset=32
                block  ;; label = @7
                  local.get 0
                  i32.const 32
                  i32.add
                  i32.const 16
                  i32.const 8
                  call $seal_get_storage
                  br_if 0 (;@7;)
                  i32.const 0
                  i32.load8_u offset=16
                  local.set 1
                end
                local.get 0
                local.get 1
                i32.const 1
                i32.and
                i32.store8 offset=31
                i32.const 1
                br_if 5 (;@1;)
                unreachable
                unreachable
              end
              unreachable
              unreachable
            end
            unreachable
            unreachable
          end
          unreachable
          unreachable
        end
        i32.const 0
        i32.const 0
        i32.const 0
        call $seal_return
        unreachable
      end
      unreachable
      unreachable
    end
    i32.const 1
    call $__malloc
    local.tee 1
    local.get 0
    i32.load8_u offset=31
    i32.store8
    i32.const 0
    local.get 1
    i32.const 1
    call $seal_return
    unreachable)
  (global $__stack_pointer (mut i32) (i32.const 65536))
  (export "deploy" (func $deploy))
  (export "call" (func $call)))
