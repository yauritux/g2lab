package com.yauritux.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.yauritux.model.entity.Card;

/**
 * 
 * @author yauritux
 * @version 1.0.0
 * @since 1.0.0
 *
 */
public interface CardRepository extends JpaRepository<Card, Long> {

	Card findBySerialNo(String serialNo);
}
